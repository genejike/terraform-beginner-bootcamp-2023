// The package main statement indicates that this Go file is part of the main package.
// In Go, the main package is special—it is the entry point for the executable programs.
package main

// This line imports the "fmt" package, which provides functions for formatted I/O, including printing to the console.

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"log"
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

// func main(): Defines the main function, the entry point of the app.
// When you run the program, it starts executing from this function.
func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: Provider,
	})
	// Format.PrintLine
	// Prints to standard output
	fmt.Println("Hello, world!")
}

type Config struct {
	Endpoint string
	Token string
	UserUuid string
}

// in golang, a titlecase function will get exported.
func Provider() *schema.Provider {
    var p *schema.Provider
    p = &schema.Provider{
        ResourcesMap: map[string]*schema.Resource{
        "terratowns_home":Resource(),
		},
        DataSourcesMap: map[string]*schema.Resource{

		},
        Schema: map[string]*schema.Schema{
			"endpoint": {
				Type: schema.TypeString,
				Required: true,
				Description: "The endpoint for hte external service",
			},
			"token": {
				Type: schema.TypeString,
				Sensitive: true, // make the token as sensitive to hide it the logs
				Required: true,
				Description: "Bearer token for authorization",
			},
			"user_uuid": {
				Type: schema.TypeString,
				Required: true,
				Description: "UUID for configuration",
				//ValidateFunc: validateUUID,
			},
		},
	}
	p.ConfigureContextFunc = providerConfigure(p)
	return p
}

func validateUUID(v interface{}, k string) (ws []string, errors []error) {
	log.Print("validateUUID:start")
	value := v.(string)
	if _,err := uuid.Parse(value); err != nil {
		errors = append(errors, fmt.Errorf("invalid UUID format"))
	}
	log.Print("validateUUID:end")
	return
}
func providerConfigure(p *schema.Provider) schema.ConfigureContextFunc{
	return func (ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		log.Print("ProviderConfigure:start")
		config:= Config {
			Endpoint:d.Get("endpoint").(string),
			Token:d.Get("token").(string),
			UserUuid:d.Get("user_uuid").(string),

		}
		log.Print("ProviderConfigure:end")
		return &config, nil
	}
}

func Resource() *schema.Resource{
	log.Print("Resource:start")

	resource:= &schema.Resource{
		CreateContext: resourceHouseCreate,
		ReadContext: resourceHouseRead,
		UpdateContext: resourceHouseUpdate,
		DeleteContext: resourceHouseDelete,
		Schema:map[string]*schema.Schema{
			"name":{
				Type: schema.TypeString,
				Required: true,
				Description:"Name of home",

			},
			"description":{
				Type: schema.TypeString,
				Required: true,
				Description:"description of home",

			},
			"domain_name":{
				Type: schema.TypeString,
				Required: true,
				Description:"Domain Name of home eg .cloudfront.net",

			},
			"town":{
				Type: schema.TypeString,
				Required: true,
				Description:"Name of town the home belongs to ",

			},
			"content_version":{
				Type: schema.TypeInt,
				Required: true,
				Description:"content_version",

			},

           
		},
	}

	log.Print("Resource:end")
    return resource
}

func resourceHouseCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Print("resourceHouseCreate:start")
	var diags diag.Diagnostics

	config := m.(*Config)

	payload := map[string]interface{}{
		"name": d.Get("name").(string),
		"description": d.Get("description").(string),
		"domain_name": d.Get("domain_name").(string),
		"town": d.Get("town").(string),
		"content_version": d.Get("content_version").(int),
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return diag.FromErr(err)
	}

	url :=  config.Endpoint+"/u/"+config.UserUuid+"/homes"
	log.Print("URL: "+ url)
	// Construct the HTTP Request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return diag.FromErr(err)
	}

	// Set Headers
	req.Header.Set("Authorization", "Bearer "+config.Token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer resp.Body.Close()

	// parse response JSON
	var responseData map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseData);  err != nil {
		return diag.FromErr(err)
	}

	// StatusOK = 200 HTTP Response Code
	if resp.StatusCode != http.StatusOK {
		return diag.FromErr(fmt.Errorf("failed to create home resource, status_code: %d, status: %s, body %s", resp.StatusCode, resp.Status, responseData))
	}

	// handle response status

	homeUUID := responseData["uuid"].(string)
	d.SetId(homeUUID)

	log.Print("resourceHouseCreate:end")

	return diags
}

func resourceHouseRead(ctx context.Context,d *schema.ResourceData,m interface{})diag.Diagnostics{
	log.Print("resourceHouseRead:start")
	var diags diag.Diagnostics
	config := m.(*Config)
    homeUUID:= d.Id()
 // Construct the Http Request 
 url := config.Endpoint+"/u/"+config.UserUuid+"/homes/"+homeUUID
	log.Print("URL: "+ url)
 req,err := http.NewRequest("GET",url,nil)
 if err !=nil{
	 return diag.FromErr(err)
 }

 //set Headers
 req.Header.Set("Authorization", "Bearer"+config.Token)
 req.Header.Set("Content-Type","application/json")
 req.Header.Set("Accept","application/json")
 client:=http.Client{}
 resp, err :=client.Do(req)
 if err !=nil{
	 return diag.FromErr(err)
	  }
 defer resp.Body.Close()
   

 var responseData map[string]interface{}
//status ok =200 Http Response code
if resp.StatusCode == http.StatusOK{
	 // parse response json 
	 var responseData map[string]interface{}
	 if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil{
		 return diag.FromErr(err)
	 }
	 d.Set("name",responseData["name"].(string))
	 d.Set("description",responseData["description"].(string))
	 d.Set("domain_name",responseData["domain_name"].(string))
	 d.Set("content_version",responseData["content_version"].(float64))
	 } else if resp.StatusCode != http.StatusOK{
		d.SetId("")
	 } else if resp.StatusCode != http.StatusOK{
	return diag.FromErr(fmt.Errorf("failed to create house resource, status_code:%d,status:%s,body %s",resp.StatusCode,resp.Status,responseData))
	
	}

	

log.Print("resourceHouseRead:end")
return diags
	
  }

func resourceHouseUpdate(ctx context.Context,d *schema.ResourceData,m interface{})diag.Diagnostics{
	log.Print("resourceHouseUpdate:start")
	var diags diag.Diagnostics
	config := m.(*Config)
    homeUUID:= d.Id()

	payload:=map[string]interface{}{
		"name": d.Get("name").(string),
		"description":d.Get("description").(string),
		"content_version":d.Get("content_version").(int),
	
		}
	
	payloadBytes, err := json.Marshal(payload)
		if err !=nil{
			return diag.FromErr(err)
		}
	
	 // Construct the Http Request 
	 url := config.Endpoint+"/u/"+config.UserUuid+"/homes/"+homeUUID
	log.Print("URL: "+ url)
	req, err := http.NewRequest("PUT",url,bytes.NewBuffer(payloadBytes))
    if err !=nil{
		return diag.FromErr(err)
	}

	//set Headers
	req.Header.Set("Authorization", "Bearer"+config.Token)
	req.Header.Set("Content-Type","application/json")
	req.Header.Set("Accept","application/json")
	client:=http.Client{}
	resp, err :=client.Do(req)
    if err !=nil{
	 return diag.FromErr(err)
		 }
	defer resp.Body.Close()

    //status ok =200 Http Response code
    if resp.StatusCode != http.StatusOK{
		return diag.FromErr(fmt.Errorf("failed to create house resource, status_code:%d,status:%s",resp.StatusCode,resp.Status))
		}
	

	log.Print("resourceHouseUpdate:end")
	d.Set("name",payload["name"])
	d.Set("description",payload["description"])
	d.Set("content_version",payload["content_version"])

	return diags
	
  }
 

func resourceHouseDelete(ctx context.Context,d *schema.ResourceData,m interface{})diag.Diagnostics{
	log.Print("resourceHouseDelete:start")
	var diags diag.Diagnostics
	config := m.(*Config)
    homeUUID:= d.Id()

	 // Construct the Http Request 
	 url :=  config.Endpoint+"/u/"+config.UserUuid+"/homes/"+homeUUID
	 log.Print("URL: "+ url)

	req,err := http.NewRequest("DELETE",url,nil)
	if err !=nil{
		return diag.FromErr(err)
	}
   
	//set Headers
	req.Header.Set("Authorization", "Bearer"+config.Token)
	req.Header.Set("Content-Type","application/json")
	req.Header.Set("Accept","application/json")
	client:=http.Client{}
	resp, err :=client.Do(req)
	if err !=nil{
	 return diag.FromErr(err)
	  }
	  defer resp.Body.Close()

    //status ok =200 Http Response code
    if resp.StatusCode != http.StatusOK{
		return diag.FromErr(fmt.Errorf("failed to create house resource, status_code:%d,status:%s",resp.StatusCode,resp.Status))
		}
	d.SetId("")
	log.Print("resourceHouseDelete:end")
	return diags
	
  }
package com.erp.app;

import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.erp.dao.DriverDAO;
import com.erp.dao.VehicleDAO;
import com.google.gson.JsonArray;
import com.google.gson.JsonObject;
import com.google.gson.JsonPrimitive;

import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.time.Duration;

import org.springframework.web.bind.annotation.GetMapping;

@RestController
@RequestMapping("/erp")
public class DriverVehicleController {

    @GetMapping("/send")
    public String SendVehicleAndDriver() {
        DriverDAO driverDAO = new DriverDAO();
        VehicleDAO vehicleDAO = new VehicleDAO();
        JsonObject postJSON = new JsonObject();
        
        JsonArray invoices = new JsonArray();
        invoices.add(new JsonPrimitive("4234545111215123123155515123112315514134412"));
        invoices.add(new JsonPrimitive("4234545111215123123155515123112315514134413"));
        invoices.add(new JsonPrimitive("4234545111215123123155515123112315514134414"));

        
        postJSON.addProperty("driverId", driverDAO.getFirstDriverAsc().getId().toString());
        postJSON.addProperty("vehiclePlate", vehicleDAO.getFirstVehicleAsc().getPlate());
        postJSON.addProperty("routeDescription", "Entrega de produtos frágeis na região central");
        postJSON.addProperty("invoices", invoices.toString());
        


        HttpClient client = HttpClient.newBuilder()
                .version(HttpClient.Version.HTTP_2) 
                .connectTimeout(Duration.ofSeconds(10))
                .build();
        
        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create("http://localhost:8081/v1/assignments/notify")) 
                .header("Content-Type", "application/json")
                .POST(HttpRequest.BodyPublishers.ofString(postJSON.toString()))
                .build();
        
        try{
            HttpResponse<String> response = client.send(request, HttpResponse.BodyHandlers.ofString());
    
            System.out.println("Status Code: " + response.statusCode());
            System.out.println("Response Body: " + response.body());
        } catch (Exception ex) {
            System.out.println(ex.getMessage());
        }


        return postJSON.toString();
    }
    
}

package com.erp.app;

import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.erp.dao.DriverDAO;
import com.erp.dao.VehicleDAO;
import com.google.gson.JsonObject;

import org.springframework.web.bind.annotation.GetMapping;

@RestController
@RequestMapping("/erp")
public class DriverVehicleController {

    @GetMapping("/send")
    public String SendVehicleAndDriver() {
        DriverDAO driverDAO = new DriverDAO();
        VehicleDAO vehicleDAO = new VehicleDAO();
        JsonObject postJSON = new JsonObject();
        
        postJSON.addProperty("driverId", driverDAO.getFirstDriverAsc().getId().toString());
        postJSON.addProperty("vehiclePlate", vehicleDAO.getFirstVehicleAsc().getPlate());
        postJSON.addProperty("routeDescription", "");
        postJSON.addProperty("invoices", "");
        

        return postJSON.toString();
    }
    
}

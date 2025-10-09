package com.erp.erp;

import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.GetMapping;


@RestController
@RequestMapping("/api")
public class ERPRestController {

    @GetMapping("/greeting")
    public String getGreeting() {
        return "Hello World";
    }
    
}

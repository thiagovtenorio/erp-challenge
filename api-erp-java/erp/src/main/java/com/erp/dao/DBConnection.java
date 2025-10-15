package com.erp.dao;

import java.io.IOException;
import java.io.InputStream;
import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.SQLException;
import java.util.Properties;

public class DBConnection {
    private String jdbcURL;
    private String username;
    private String password;
    
    public DBConnection() {
        Properties properties = new Properties();

        try (InputStream input = DBConnection.class.getClassLoader().getResourceAsStream("config.properties")) {
            properties.load(input);
    
            this.jdbcURL = properties.getProperty("database.url");
            this.username = properties.getProperty("database.username");
            this.password = properties.getProperty("database.password");
    
            } catch (IOException ex) {
                ex.printStackTrace();
            }
    
    }
    
    public Connection getConnection() throws SQLException{
        return DriverManager.getConnection(this.jdbcURL, this.username, this.password);
    }
}

package com.erp.dao;

import java.sql.Connection;
import java.sql.ResultSet;
import java.sql.Statement;
import java.util.UUID;

import com.erp.model.Driver;

public class DriverDAO extends DAO{

    public DriverDAO(){
        super();
    }

    public Driver getFirstDriverAsc (){
        Driver driver = null;
        try {
            conn = getConnection();
            statement = conn.createStatement();
            sql = "select d.id as driver_id, d.created_at from drivers d where d.status = 'AVAILABLE' order by d.created_at asc limit 1;";
            rs = statement.executeQuery(sql);
            
            if(rs.next()){
                driver = new Driver();
                driver.setId((UUID) rs.getObject("driver_id"));
            }
        } catch (Exception ex){

        } finally {

        }

        return driver;
    }

}

package com.erp.dao;

import com.erp.model.Vehicle;

public class VehicleDAO extends DAO{
    public VehicleDAO(){
        super();
    }

    public Vehicle getFirstVehicleAsc(){
        Vehicle vehicle = null;
        try{
            conn = getConnection();
            statement = conn.createStatement();
            sql = "select v.plate_number as vehicle_plate from vehicles v where  v.status = 'OPERATIONAL' order by v.created_at asc limit 1;";
            rs = statement.executeQuery(sql);

            if(rs.next()){
                vehicle = new Vehicle();
                vehicle.setPlate(rs.getString("vehicle_plate"));
            }
        }catch(Exception ex){

        }
        return vehicle;
    }
}

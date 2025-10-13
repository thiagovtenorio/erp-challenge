package com.erp.dao;

import java.sql.Connection;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.sql.Statement;

public class DAO {
    private DBConnection dbConnection;
    protected Connection conn;
    protected Statement statement;
    protected String sql;
    protected ResultSet rs;

    public DAO(){
        try{
            this.dbConnection = new DBConnection();
        } catch(Exception ex){
            System.out.println(ex.getMessage());
        }
    }
    public Connection getConnection () throws SQLException{
        return dbConnection.getConnection();
    }
}

Test for  Sim Sim

Configuration
--------------

Configuration file contains database related information 

{
  "user": "root",
  "dbname": "simsimtest",
  "password": "Scala@1234",
  "dbaddress": "tcp(127.0.0.1:3306)",
  "generate": "true"
}

Api Service
-----------

this service use echo web api framework and use simplest way to get data for shift swap functionality.

Contract
-------
This directory consits of the all the busines related contract objects.

Config
------

This directory contains functions to read configuration file. 

Business
--------

this contains all the functionality for manipulating business rules.

Repository
----------

This folder contains functions to read data from GORM orm and send back to business rules engine.


 Entities
 --------
 
 This folder contains data structure for tables and schema mapping for the database.
 
 Connection
 ----------
 
 Connection will create the connection against the database if database does nt exits it will create database in local mysql server based on given ip path on configuration file.
 
 Error Handle
 ------------
 
 Handle error and recover panic if service crash occur.
 
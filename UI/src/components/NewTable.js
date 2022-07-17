import React from "react";
import Table from 'react-bootstrap/Table';
import 'bootstrap/dist/css/bootstrap.min.css';
import { NavItem } from "react-bootstrap";


function NewTable(table){

    return(
    <Table striped bordered hover>
        table.map((item)=>(
          <tr>
            <td>{item.Titles}</td>
          </tr>
        ))
       
        <tr>
          <td>col1</td>
          <td>col1</td>
          <td>col1</td>
          <td>col1</td>
        </tr>
        <tr>
          <td>col1</td>
          <td>col1</td>
          <td>col1</td>
          <td>col1</td>
        </tr>
        <tr>
          <td>col1</td>
          <td>col1</td>
          <td>col1</td>
          <td>col1</td>
        </tr>
        
      </Table>)
}

export default NewTable;
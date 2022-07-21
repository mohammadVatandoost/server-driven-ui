import React from "react";
import Table from 'react-bootstrap/Table';
import 'bootstrap/dist/css/bootstrap.min.css';


function NewTable(table){
    
    const Titles = () => {
      return table.table.Titles.map( (item)=> {
        return (
       
          <th>{item}</th>
       
        )
      })
    }

    const Rows = () => {
      return table.table.Rows.map( (items)=> {
        let tmp = items.map((item)=>{
          return (
               <td>{item}</td> 
            )
          })
        return (
          <tr> {tmp}</tr>
        ) 
       
      })
    }

    return(
    <Table striped bordered hover>
      <tr>{Titles()}</tr>
      {Rows()}
    </Table>)
}

export default NewTable;
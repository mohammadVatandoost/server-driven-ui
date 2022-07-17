import React,{useState,useEffect} from 'react';
import logo from './logo.svg';
import './App.css';
import axios from 'axios';
import 'bootstrap/dist/css/bootstrap.min.css';
import NewTable from './components/NewTable'
import ImageComponent from './components/ImageComponent'
import { Table } from 'react-bootstrap';

function App() {
  const [data, setData] = useState();
  const [output, setOutput] = useState();

  

  useEffect(() => {
    const fetchData = async () => {
      const results = await axios(
        '/v1/page/demo',
      );
      console.log('data:',results)
      let result = []
      results.data.payload.map((item) => {
        if(item.Type == "Title") {
          result.push(<h1>{item.Data.Title}</h1>)
        } 
        else if(item.Type == "Image") {
          result.push(<img src={item.Data.Source} />)
        } 
        else if(item.Type == "Description") {
          result.push(<p>{item.Data.Description}</p>)
        } 
        else if(item.Type == "Title") {
          result.push(<NewTable table={item.Data.Titles}/>)
        } 
        
      })
      setOutput(result)
      console.log('result:',result)
      // setData(results.data);
      // const imageData =data.payload[2]
      // console.log(imageData)
    };

    fetchData();
    

  }, []);
 


  return (
   <div className='container'>
    {output}
    {/* <h1>Demo Page</h1>
    <ImageComponent />
    <div className='dataList'>
      <h2></h2>
      <NewTable/>
    </div> */}
   </div>
  );
}

export default App;

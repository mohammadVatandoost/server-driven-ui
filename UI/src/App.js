import React,{useState,useEffect} from 'react';
import './App.css';
import axios from 'axios';
import 'bootstrap/dist/css/bootstrap.min.css';
import NewTable from './components/NewTable'

function App() {
  const [output, setOutput] = useState();

  useEffect(() => {
    const fetchData = async () => {
      const results = await axios(
        '/v1/page/demo',
      );
      console.log('data:',results)
      let result = []
      results.data.payload.map((item) => {
        if(item.Type === "Title") {
          result.push(<h1>{item.Data.Title}</h1>)
        } 
        else if(item.Type === "Image") {
          result.push(<img src={item.Data.Source} />)
        } 
        else if(item.Type === "Description") {
          result.push(<p>{item.Data.Description}</p>)
        } 
        else if(item.Type === "Table") {
          result.push(<NewTable table={item.Data}/>)
        } 
        
      })
      setOutput(result)
    };

    fetchData();
    
  }, []);

  return (
   <div className='container'>
    {output}
   </div>
  );
}

export default App;

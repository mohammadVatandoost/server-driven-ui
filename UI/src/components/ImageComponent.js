import React from "react";



function ImageComponent(image){

    return(
        <div className='image'>
      <img src={image.data} alt='first image'/>
    </div>
   )
}

export default ImageComponent;
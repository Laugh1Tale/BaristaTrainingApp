import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Link, Routes, Route } from "react-router-dom"
import 'materialize-css';
import { InformationPage } from './InformationPage'
import '../css/courses.css';

export const LectionPage = () => {
    const [lections, setLection] = useState([]);

    useEffect(() => {
      const fetchData = async () => {
        try {
          const response = await axios.get('http://localhost:8000/api/lections');
          setLection(response.data);
        } catch (error) {
          console.error('Error fetching data: ', error);
        }
      };
  
      fetchData();
    }, []);

    let my_href="/lections"

    return (
      <div class="ag-format-container">
      <div class="ag-courses_box">

        {lections.map((item, index) => (

          <div class="ag-courses_item">
            <p style={{display: 'none'}}>{index === 0 ? my_href = "/information" : my_href = "/lections"}</p>
            <a href={my_href} class="ag-courses-item_link">
            <div class="ag-courses-item_bg"></div>

            <div class="ag-courses-item_title">
              {item.theme}
            </div>

          </a>
        </div>
        ))}

      </div>
    </div>
    )
}
import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Link, Routes, Route } from "react-router-dom"
import 'materialize-css';
import { LessonsPage } from './LessonsPage'
import '../css/courses.css';

export const CoursesPage = () => {
    const [courses, setCourse] = useState([]);

    useEffect(() => {
      const fetchData = async () => {
        try {
          const response = await axios.get('http://localhost:8000/api/courses'); 
          setCourse(response.data);
        } catch (error) {
          console.error('Error fetching data: ', error);
        }
      };
  
      fetchData();
    }, []);

    let my_href= "/courses";

    return (
      <div class="ag-format-container">
        <div class="ag-courses_box">

          {courses.map((item, index) => (
            <div class="ag-courses_item">
              <p style={{display: 'none'}}>{index === 0 ? my_href = "/lessons" : my_href = "/courses"}</p>
              <a href={my_href} className="ag-courses-item_link">
              <div class="ag-courses-item_bg"></div>

              <div class="ag-courses-item_title">
                {item.name}
              </div>

              <div class="ag-courses-item_date-box">
                {item.description}
              </div>

              <div class="ag-courses-item_date-box">
                {"\n"}
              </div>

            <div class="ag-courses-item_date-box">
                Pasing Score:
              <span class="ag-courses-item_date">
                {item.passing_score}
              </span>
            </div>
            </a>
          </div>
          ))}

        </div>
      </div>
    )
}
import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Link, Routes, Route } from "react-router-dom"
import 'materialize-css';
import { LectionPage } from './LectionPage'
import { TestPoints } from './QuestionsPage'
import '../css/courses.css';

export const LessonsPage = () => {
    const [lessons, setLesson] = useState([]);

    useEffect(() => {
      const fetchData = async () => {
        try {
          const response = await axios.get('http://localhost:8000/api/lessons');
          setLesson(response.data);
        } catch (error) {
          console.error('Error fetching data: ', error);
        }
      };
  
      fetchData();
    }, []);

    let my_href="/lessons"

    return (
      <div class="ag-format-container">
      <div class="ag-courses_box">

        {lessons.map((item, index) => (

          <div class="ag-courses_item">
            {index === 0 ? my_href = "/lections" : my_href = "/lessons"}
            <a href={my_href} class="ag-courses-item_link">
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
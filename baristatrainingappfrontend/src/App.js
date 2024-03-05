import React from 'react';
import 'materialize-css';
import {BrowserRouter as Router, Routes, Navigate, Route, redirect} from 'react-router-dom'
import {CoursesPage} from './pages/CoursesPage'
import {LessonsPage} from './pages/LessonsPage'
import {QuestionsPage} from './pages/QuestionsPage'
import {TestsPage} from './pages/TestsPage'
import { SignIn } from './pages/SignIn'
import { SignUp } from './pages/SignUp'
import { LectionPage } from './pages/LectionPage';
import { InformationPage } from './pages/InformationPage';

var TestPoints=0

function App() {
  const isAuthentificated=true
  if (isAuthentificated) {
    return (
      <Router>
        <Routes>
          <Route path="/*" element={<CoursesPage />} />
          <Route path="/courses" element={<CoursesPage />} />
          <Route path="/lessons" element={<LessonsPage />} />
          <Route path="/lections" element={<LectionPage />} />
          <Route path="/information" element={<InformationPage />} />
          {/* <Route path="/questions" element={<QuestionsPage />} /> */}
          <Route path="/tests" element={<TestsPage />} />
        </Routes>
      </Router>
    );
  }
  return (
    <Router>
      <Routes>
        <Route path="/*" element={<SignIn />} />
        <Route path="/sign-in" element={<SignIn />} />
        <Route path="/sign-up" element={<SignUp />} />
      </Routes>
    </Router>   
  );
}

export default App;

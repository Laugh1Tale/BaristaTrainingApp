import React from 'react';
import 'materialize-css';
import {BrowserRouter as Router, Routes, Navigate, Route, redirect} from 'react-router-dom'
import {CoursesPage} from './pages/CoursesPage'
import {LessonsPage} from './pages/LessonsPage'
import {TestsPage} from './pages/TestsPage'
import { SignIn } from './pages/SignIn'
import { SignUp } from './pages/SignUp'
import { LectionPage } from './pages/LectionPage';
import { InformationPage } from './pages/InformationPage';
import { useAuth } from './hooks/auth.hook';

var TestPoints=0

function App() {
  //const {token, login, logout, userId} = useAuth()
  const isAuthentificated = true
  if (isAuthentificated) {
    return (
      <Router>
        <Routes>
          <Route path="/*" element={<SignIn />} />
          <Route path="/sign-in" element={<SignIn />} />
          <Route path="/sign-up" element={<SignUp />} />
          <Route path="/courses" element={<CoursesPage />} />
          <Route path="/lessons" element={<LessonsPage />} />
          <Route path="/lections" element={<LectionPage />} />
          <Route path="/information" element={<InformationPage />} />
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

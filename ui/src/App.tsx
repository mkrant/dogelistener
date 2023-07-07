import React, {useState} from 'react';
import './App.css';
import '@fontsource/roboto/300.css';
import '@fontsource/roboto/400.css';
import '@fontsource/roboto/500.css';
import '@fontsource/roboto/700.css';
import { QueryClient, QueryClientProvider, useQuery } from 'react-query'
import {Sessions} from "./pages/sessions";
import { ThemeProvider, createTheme } from '@mui/material/styles';
import CssBaseline from '@mui/material/CssBaseline';
import { BrowserRouter as Router, Route, Routes, Link } from "react-router-dom";
import {SessionDetails} from "./pages/sessionDetails";


const darkTheme = createTheme({
    palette: {
        mode: 'dark',
    },
});

const queryClient = new QueryClient()

export interface IAppProps {}

const App: React.FC<IAppProps> = (props) => {
    return (

        <div>
        {/*// <ThemeProvider theme={darkTheme}>*/}
                <CssBaseline />
                <QueryClientProvider client={queryClient} contextSharing={true}>
                    <Router>
                        <Routes>
                            <Route path="/" element={<Sessions/>}/>
                            <Route path="/sessions/:id" element={<SessionDetails/>}/>
                        </Routes>
                    </Router>
                </QueryClientProvider>
        {/*// </ThemeProvider>*/}
        </div>

  );
}

export default App;

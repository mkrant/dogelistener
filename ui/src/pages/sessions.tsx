import React, {useState} from "react";
import dayjs, {Dayjs} from "dayjs";
import {useQuery} from "react-query";
import {getSessions} from "../api/api";
import {LocalizationProvider} from "@mui/x-date-pickers/LocalizationProvider";
import {AdapterDayjs} from "@mui/x-date-pickers/AdapterDayjs";
import {DatePicker} from "@mui/x-date-pickers/DatePicker";
import TextField from "@mui/material/TextField";
import {Button, Container} from "@mui/material";
import {SessionTable} from "../components/sessionTable";

export const Sessions = () => {
    const [date, setDate] = useState<Dayjs | null>(dayjs());
    const {data: sessions} = useQuery("querykey", () => {
        const resp = getSessions();
        console.log(resp);

       return resp;
    })

    return (
        <Container className="App" sx={{ m: 3 }}>
            {sessions && <SessionTable sessions={sessions.data}/>}
                
        </Container>
    )
}
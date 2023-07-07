import React, {useEffect, useState} from "react";
import {useQuery} from "react-query";
import {getRunDetail, getSession} from "../api/api";
import {Container} from "@mui/material";
import {useParams} from "react-router-dom";
import {RunTable} from "../components/runTable";
import {RunDetail} from "../types/dogelistener";
import {RunGraph} from "../components/runGraph";

export const SessionDetails = () => {
    const { id } = useParams();
    const [currentRun, setCurrentRun] = useState<RunDetail>({
        data: [],
        duration_seconds: 0,
        id: "",
        live: false,
        start_time: ""
    })

    const {data: session} = useQuery("sess" + id, {
        queryFn: () => {
                return getSession(id!);
            }
    })

    const {data: queryRun} = useQuery(["runs", session], {
        queryFn: () => {
            if (!session) {
                return null;
            }

            // if (!session.data.is_running) {
            //     return null;
            // }

            return getRunDetail(session.data.id, session.data.runs[0].id)
        }
    })

    // useEffect(() => {
    //     if (session?.data.is_running) {
    //         console.log("it running")
    //         getRunDetail(session.data.id, session.data.runs[0].id).then(d => {
    //             setCurrentRun(d.data)
    //         })
    //     } else {
    //         console.log("not running")
    //     }
    // })

    return (
        <Container className="App" sx={{ m: 3 }}>
            <h3>{"Session #" + id}</h3>
            {queryRun && queryRun.data && queryRun.data.data.length > 0 && <div>

                <RunGraph  data={queryRun.data.data}/>
            </div>}
            {session && <RunTable sessionDetail={session.data}/>}
        </Container>
    )
}
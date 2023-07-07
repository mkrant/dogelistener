import React, {FC, FunctionComponent, useEffect} from "react";
import {DogeSession, Run, SessionDetail} from "../types/dogelistener";
import {Button, Chip, Container} from "@mui/material";
import {Link} from "react-router-dom";

type RunTableProps = {
    sessionDetail: SessionDetail
}

export const RunTable: FC<RunTableProps> = ({sessionDetail}: RunTableProps) => {
    return (
        <Container>
            {sessionDetail && sessionDetail.runs.map((run) => {
                return <Chip sx={{ m: 0.5 }} color={run.live ? "success" : "primary"} variant="outlined" label={run.id + " - " + run.start_time}></Chip>
            })}
        </Container>
    )
}



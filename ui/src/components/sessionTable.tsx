import React, {FC, FunctionComponent} from "react";
import {DogeSession} from "../types/dogelistener";
import {Button, Chip, Container} from "@mui/material";
import {Link} from "react-router-dom";

type SessionTableProps = {
    sessions: DogeSession[]
}

export const SessionTable: FC<SessionTableProps> = ({sessions}: SessionTableProps) => {
    return (
        <Container>
            <h3>Active Sessions</h3>
            {sessions.map((sess) => {
                return <Link to={"/sessions/" + sess.id} key={sess.id}>
                    <Chip sx={{ m: 0.5 }} color={sess.is_running ? "success" : "primary"} variant="outlined" label={sess.id}></Chip>
                </Link>
            })}
        </Container>
    )
}



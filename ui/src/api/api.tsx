import ky from "ky";
import {RunDetail, DogeSession, SessionDetail} from "../types/dogelistener";

export const getSessions = async () => {
    return await ky.get(`http://localhost:8080/api/sessions`).json<{data: DogeSession[]}>();
}

export const getSession = async (id: string) => {
    return await ky.get(`http://localhost:8080/api/sessions/${id}`).json<{data: SessionDetail}>();
}

export const startRun = async (id: string) => {
    return await ky.get(`http://localhost:8080/api/sessions/${id}/start`).json<{}>();
}

export const stopRun = async (id: string) => {
    return await ky.get(`http://localhost:8080/api/sessions/${id}/start`).json<{}>();
}

export const getRunDetail = async (id: string, runID: string) => {
    return await ky.get(`http://localhost:8080/api/sessions/${id}/runs/${runID}`).json<{data: RunDetail}>();
}


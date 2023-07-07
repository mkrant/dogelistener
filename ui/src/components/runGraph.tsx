import React, {FC, FunctionComponent, useEffect} from "react";
import {DogeSession, Run, SessionDetail} from "../types/dogelistener";
import {Button, Chip, Container} from "@mui/material";
import {Link} from "react-router-dom";
import {Area, AreaChart, CartesianGrid, Line, LineChart, Tooltip, XAxis, YAxis} from "recharts";

type RunGraphProps = {
    data: number[];
}

type Value = {
    second: number,
    value: number,
}

export const RunGraph: FC<RunGraphProps> = ({data}: RunGraphProps) => {
    // const data = [
    //     { second: 'Page A', uv: 1000, pv: 2400, amt: 2400, uvError: [75, 20] },
    //     { second: 'Page B', uv: 300, pv: 4567, amt: 2400, uvError: [90, 40] },
    //     { second: 'Page C', uv: 280, pv: 1398, amt: 2400, uvError: 40 },
    //     { second: 'Page D', uv: 200, pv: 9800, amt: 2400, uvError: 20 },
    //     { second: 'Page E', uv: 278, pv: null, amt: 2400, uvError: 28 },
    //     { name: 'Page F', uv: 189, pv: 4800, amt: 2400, uvError: [90, 20] },
    //     { name: 'Page G', uv: 189, pv: 4800, amt: 2400, uvError: [28, 40] },
    //     { name: 'Page H', uv: 189, pv: 4800, amt: 2400, uvError: 28 },
    //     { name: 'Page I', uv: 189, pv: 4800, amt: 2400, uvError: 28 },
    //     { name: 'Page J', uv: 189, pv: 4800, amt: 2400, uvError: [15, 60] },
    // ];

    let dataValues: Value[] = []
    let idx = 0;
    data.map(value => {
        dataValues.push({
            second: idx,
            value: value,
        })
        idx++;
    })

    return (
<div>
        <LineChart
            width={1600}
            height={500}
            data={dataValues}
            margin={{ top: 5, right: 20, left: 10, bottom: 5 }}
        >
            <XAxis dataKey="second" />
            <YAxis/>
            <Tooltip />
            <CartesianGrid stroke="#eee"  fill="#e4ecf7" verticalPoints={[1,2,3,4]}  />
            <Line type="monotone" dataKey="value" stroke="#8884d8"  yAxisId={0} dot={false} strokeWidth={2} />
        </LineChart>

    <AreaChart          width={1600}
                        height={500}
                        data={dataValues}
                        margin={{ top: 5, right: 20, left: 10, bottom: 5 }}>
        <defs>
            <linearGradient id="colorUv" x1="0" y1="0" x2="0" y2="1">
                <stop offset="5%" stopColor="#82ca9d" stopOpacity={1}/>
                <stop offset="95%" stopColor="#82ca9d" stopOpacity={0.6}/>
            </linearGradient>
        </defs>
        <XAxis dataKey="second" />
        <YAxis />
        <CartesianGrid strokeDasharray="3 3" />
        <Tooltip />
        <Area type="monotone" dataKey="value" stroke="#82ca9d" fillOpacity={1} fill="url(#colorUv)" />
    </AreaChart>
</div>
    )
}



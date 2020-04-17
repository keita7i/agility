import React, { useEffect, useState } from 'react';

function App() {
        let [sprints, setSprints] = useState([]);

        function loadSprints(f) {
                fetch('/v1/sprints').then((res) => {
                        res.json().then((ss) => {
                                f(ss);
                        });
                });
        }

        useEffect(() => {
                loadSprints((ss) => {
                        setSprints(ss);
                });
        }, []);

        function metricTable(team) {
                return (
                        <div>
                                <style jsx>{`
                                        h2 {
                                                color: #666;
                                                font-size: large;
                                        }

                                        table {
                                                border-width: 1pt 0 0 1pt;
                                                border-color: #ccc;
                                                border-style: solid;
                                                border-collapse: collapse;
                                                margin-bottom: 6ex;
                                        }
                                `}</style>
                                <h2>{team}</h2>
                                <table>
                                        <thead>
                                                {headerRow()}
                                        </thead>
                                        <tbody>
                                                {metricRow(team, 'Commitment')}
                                                {metricRow(team, 'Done')}
                                                {metricRow(team, 'Velocity')}
                                        </tbody>
                                </table>
                        </div>
                );
        }

        function headerRow() {
                return (<tr>
                        <style jsx>{`
                                th {
                                        border-width: 0 1pt 1pt 0;
                                        border-color: #ccc;
                                        border-style: solid;
                                        padding: 0.5ex;
                                        text-align: left;
                                }

                                .metric-category {
                                        width: 8em;
                                }

                                .sprint {
                                        width: 5em;

                                }
                        `}</style>
                        <th className="metric-category"></th>
                        {sprints.map((s, i) => {
                                return <th className="sprint" key={i}>{s.sprint}</th>;
                        })}
                </tr>);
        }

        function metricRow(team, category) {
                return (<tr>
                        <style jsx>{`
                                th {
                                        border-width: 0 1pt 1pt 0;
                                        border-color: #ccc;
                                        border-style: solid;
                                        padding: 0.5ex;
                                        text-align: left;
                                }

                                td {
                                        border-width: 0 1pt 1pt 0;
                                        border-color: #ccc;
                                        border-style: solid;
                                        padding: 0.5ex;
                                        text-align: right;
                                }
                        `}</style>
                        <th>{category}</th>
                        {sprints.map((s, i) => {
                                return <td key={i}>{s.teams[team][category.toLowerCase()]}</td>;
                        })}
                </tr>);
        }

        return (
                <div className="app">
                        <style jsx>{`
                                .tables {
                                        margin: 1rem auto;
                                        width: 800pt;
                                }

                                h1 {
                                        background: #066;
                                        margin: 0 0 2rem;
                                        padding: 1rem;
                                        color: #fff;
                                        font-size: large;
                                        font-weight: normal;
                                }
                        `}</style>
                        <h1>Agility</h1>
                        <div className="tables">
                                {metricTable('All')}
                                {metricTable('SRE 0')}
                                {metricTable('SRE 1+2')}
                        </div>
                </div>
        );
}

export default App;

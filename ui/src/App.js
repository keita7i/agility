import React, { useEffect, useState } from 'react';

function App() {
        let [teams, setTeams] = useState([]);
        let [sprints, setSprints] = useState([]);

        function loadTeams(f) {
                fetch('/v1/teams').then(res => {
                        res.json().then(ts => {
                                f(ts);
                        });
                })
        }

        function loadSprints(f) {
                fetch('/v1/sprints').then((res) => {
                        res.json().then((ss) => {
                                ss.reverse();
                                f(ss);
                        });
                });
        }

        useEffect(() => {
                loadTeams(ts => {
                        setTeams(ts);
                });
                loadSprints(ss => {
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
                                return <th className="sprint" key={i}>s{s.sprint}</th>;
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
                                let metric = s.teams[team][category.toLowerCase()];
                                return <td key={i}>{metric >= 0 ? metric : null}</td>;
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
                                {teams.map(t => metricTable(t))}
                        </div>
                </div>
        );
}

export default App;

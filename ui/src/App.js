import React, { useEffect, useState } from 'react';

const RELOADING_INTERVAL_MILLIS = 5000;

function App() {
        const [boards, setBoards] = useState([]);

        function loadBoards(f) {
                fetch('/v1/boards').then((res) => {
                        res.json().then((ss) => {
                                f(ss);
                        });
                });
        }

        useEffect(() => {
                loadBoards(boards => {
                        setBoards(boards);
                });
        }, []);

        function teamBoard(board, key) {
                const sprints = [...board.sprints];
                sprints.reverse();

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
                                <th className="metric-category">Sprint</th>
                                {sprints.map((s, i) => {
                                        return <th className="sprint" key={i}>{s.name}</th>;
                                })}
                        </tr>);
                }

                function metricRow(category) {
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
                                        let metric = s[category.toLowerCase()];
                                        return <td key={i}>{metric >= 0 ? metric : null}</td>;
                                })}
                        </tr>);
                }

                return (
                        <div key={key}>
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
                                <h2>{board.team}</h2>
                                <table>
                                        <thead>
                                                {headerRow()}
                                        </thead>
                                        <tbody>
                                                {metricRow('Commitment')}
                                                {metricRow('Velocity')}
                                        </tbody>
                                </table>
                        </div>
                );
        }

        return (
                <div className="app">
                        <style jsx>{`
                                .tables-wrapper {
                                        position: relative;
                                        width: 100%;
                                        padding: 2rem;
                                        background: ${boards.length <= 0 ? "gray" : "none"};
                                        opacity: ${boards.length <= 0 ? "0.5" : "1"};
                                }

                                .tables {
                                        position: relative;
                                        margin: 0 auto;
                                        width: 800pt;
                                        z-index: -99;
                                }

                                h1 {
                                        background: #066;
                                        margin: 0;
                                        padding: 1rem;
                                        color: #fff;
                                        font-size: large;
                                        font-weight: normal;
                                }
                        `}</style>
                        <h1>Agility</h1>
                        <div className="tables-wrapper">
                                <div className="tables">
                                        {boards.map(board => teamBoard(board))}
                                </div>
                        </div>
                </div>
        );
}

export default App;

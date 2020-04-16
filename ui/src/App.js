import React from 'react';

function App() {
        return (
                <div className="app">
                        <style jsx>{`
                                .app {
                                        margin: 1rem auto;
                                        width: 900pt;
                                }

                                h1 {
                                        background: #066;
                                        margin: 0 0 2rem;
                                        padding: 1rem;
                                        color: #fff;
                                        font-size: large;
                                        font-weight: normal;
                                }

                                h2 {
                                        color: #666;
                                        font-size: large;
                                }

                                table {
                                        border-width: 1pt 0 0 1pt;
                                        border-color: #ccc;
                                        border-style: solid;
                                        border-collapse: collapse;
                                }

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

                                .metric-category {
                                        width: 8em;
                                }

                                .sprint {
                                        width: 5em;

                                }

                                .metric-table {
                                        margin-bottom: 6ex;
                                }
                        `}</style>
                        <h1>Agility</h1>
                        <h2>All</h2>
                        <table className="metric-table">
                                <tr>
                                        <th className="metric-category"></th>
                                        <th className="sprint">s11</th>
                                        <th className="sprint">s12</th>
                                </tr>
                                <tr>
                                        <th>Commitment</th>
                                        <td></td>
                                        <td></td>
                                </tr>
                                <tr>
                                        <th>Done</th>
                                        <td></td>
                                        <td></td>
                                </tr>
                                <tr>
                                        <th>Velocity</th>
                                        <td></td>
                                        <td></td>
                                </tr>
                        </table>
                        <h2>SRE 0</h2>
                        <table className="metric-table">
                                <tr>
                                        <th className="metric-category"></th>
                                        <th className="sprint">s11</th>
                                        <th className="sprint">s12</th>
                                </tr>
                                <tr>
                                        <th>Commitment</th>
                                        <td>20 (5)</td>
                                        <td></td>
                                </tr>
                                <tr>
                                        <th>Done</th>
                                        <td></td>
                                        <td></td>
                                </tr>
                                <tr>
                                        <th>Velocity</th>
                                        <td></td>
                                        <td></td>
                                </tr>
                        </table>
                        <h2>SRE 1+2</h2>
                        <table className="metric-table">
                                <tr>
                                        <th className="metric-category"></th>
                                        <th className="sprint">s11</th>
                                        <th className="sprint">s12</th>
                                </tr>
                                <tr>
                                        <th>Commitment</th>
                                        <td></td>
                                        <td></td>
                                </tr>
                                <tr>
                                        <th>Done</th>
                                        <td></td>
                                        <td></td>
                                </tr>
                                <tr>
                                        <th>Velocity</th>
                                        <td></td>
                                        <td></td>
                                </tr>
                        </table>
                </div>
        );
}

export default App;

import React from 'react'
import Plotly from 'react-plotlyjs'

const trace1 = {
  x: [1, 2, 3, 4],
  y: [10, 15, 13, 17],
  mode: 'markers',
  type: 'scatter'
}

const trace2 = {
  x: [2, 3, 4, 5],
  y: [16, 5, 11, 9],
  mode: 'lines',
  type: 'scatter'
}

const trace3 = {
  x: [1, 2, 3, 4],
  y: [12, 9, 15, 12],
  mode: 'lines+markers',
  type: 'scatter'
}

const data = [trace1, trace2, trace3]
const layout = {
  font: {
    family: 'Courier New, monospace',
    size: 18,
    color: '#fff'
  },
  paper_bgcolor: 'rgba(0,0,0,0)',
  plot_bgcolor: 'rgba(0,0,0,0)'
}

const StatsView = () => (
  <div>
  <section className="page-header">
    <h1>
      Stats
    </h1>
  </section>
  <section class="row">
    <Plotly data={data} layout={layout}/>
    <Plotly data={data} layout={layout}/>
    <Plotly data={data} layout={layout}/>
  </section>
  </div>
)

export default StatsView

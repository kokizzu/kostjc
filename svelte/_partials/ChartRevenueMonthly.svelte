<script>
  /** @typedef {import('../_types/property.js').ChartRevenueReport} ChartRevenueReport */

  import Chart from 'chart.js/auto';  

  let chart = /** @type {import('chart.js').Chart} */ (null);

  export let chartRevenueReports = /** @type {ChartRevenueReport[]} */ ([]);

  function buildDatasets(data) {
    return [
      {
        label: 'Cash',
        type: 'bar',
        data: (data || []).map((/** @type {ChartRevenueReport} */ i) => Number(i.cashIDR || 0)),
        borderColor: '#22c55e',
        backgroundColor: '#22c55e60',
        borderWidth: 1,
        stack: 'revenue',
      },
      {
        label: 'Transfer',
        type: 'bar',
        data: (data || []).map((/** @type {ChartRevenueReport} */ i) => Number(i.transferIDR || 0)),
        borderColor: '#38bdf8',
        backgroundColor: '#38bdf860',
        borderWidth: 1,
        stack: 'revenue',
      },
      {
        label: 'Total',
        type: 'line',
        data: (data || []).map((/** @type {ChartRevenueReport} */ i) => Number(i.revenueIDR || 0)),
        borderColor: '#2563eb',
        backgroundColor: '#2563eb20',
        pointRadius: 0,
        tension: 0.1,
        fill: false,
        pointStyle: 'circle',
        pointBackgroundColor: '#2563eb',
        pointBorderWidth: 2,
        pointHitRadius: 10
      },
      {
        label: 'Donation',
        type: 'bar',
        data: (data || []).map((/** @type {ChartRevenueReport} */ i) => Number(i.donationIDR || 0)),
        borderColor: '#f59e0b',
        backgroundColor: '#f59e0b60',
        borderWidth: 1,
        stack: 'revenue',
      }
    ];
  }

  setTimeout(() => {
    const ElmChart = /** @type {HTMLCanvasElement} */ (document.getElementById('chart'));
    chart = new Chart(ElmChart, {
      type: 'line',
      data: {
        labels: (chartRevenueReports || []).map((/** @type {ChartRevenueReport} */ i) => {
          const dt = /** @type {Date} */ (new Date(i.date));
          return dt.toLocaleDateString('en-US', {
            month: 'short',
            day: '2-digit'
          });
        }),
        datasets: buildDatasets(chartRevenueReports)
      },
      options: {
        plugins: {
          legend: {
            display: true
          },
          tooltip: {
            enabled: true
          }
        },
        maintainAspectRatio: false,
        responsive: true,
        scales: {
          x: {
            stacked: true
          },
          y: {
            beginAtZero: true,
            stacked: true
          }
        }
      }
    });
  }, 400);

  export function updateData(/** @type {ChartRevenueReport[]} */ data) {
    if (!chart) {
      return
    }

    chart.data.labels = (data || []).map((/** @type {ChartRevenueReport} */ i) => {
      const dt = /** @type {Date} */ (new Date(i.date));
      return dt.toLocaleDateString('en-US', {
        month: 'short',
        day: '2-digit'
      });
    }),
    chart.data.datasets = buildDatasets(data);
    
    chart.update();
  }
</script>

<div class="chart">
  <canvas id="chart"></canvas>
</div>

<style>
  .chart {
    height: 400px;
    width: 100%;
    display: flex;
    width: 100%;
    padding: 10px;
    border: 1px solid var(--gray-003);
    border-radius: 8px;
  }

  .chart canvas {
    width: 100% !important;
    height: 100% !important;
  }
</style>

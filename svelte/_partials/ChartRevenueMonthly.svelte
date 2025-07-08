<script>
  /** @typedef {import('../_types/property.js').ChartRevenueReport} ChartRevenueReport */

  import Chart from 'chart.js/auto';  

  let chart = /** @type {import('chart.js').Chart} */ (null);

  export let chartRevenueReports = /** @type {ChartRevenueReport[]} */ ([]);

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
        datasets: [{
          label: 'Revenue',
          data: (chartRevenueReports || []).map((/** @type {ChartRevenueReport} */ i) => Number(i.revenueIDR)),
          borderColor: '#2563eb',
          backgroundColor: '#2563eb60',
          pointRadius: 0,
          tension: 0.1,
          fill: true,
          pointStyle: 'circle',
          pointBackgroundColor: '#2563eb',
          pointBorderWidth: 2,
          pointHitRadius: 10
        }]
      },
      options: {
        plugins: {
          legend: {
            display: false
          },
          tooltip: {
            enabled: true
          }
        },
        maintainAspectRatio: false,
        responsive: true,
        scales: {
          y: {
            beginAtZero: true
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
    chart.data.datasets = [{
      label: 'Revenue',
      data: (data || []).map((/** @type {ChartRevenueReport} */ i) => Number(i.revenueIDR)),
      borderColor: '#2563eb',
      backgroundColor: '#2563eb60',
      pointRadius: 0,
      tension: 0.1,
      fill: true,
      pointStyle: 'circle',
      pointBackgroundColor: '#2563eb',
      pointBorderWidth: 2,
      pointHitRadius: 10
    }];
    
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

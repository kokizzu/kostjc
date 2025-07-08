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
          borderColor: '#f97316',
          backgroundColor: '#f9731630',
          pointRadius: 0,
          tension: 0.1
        }]
      },
      options: {
        plugins: {
          legend: {
            display: false
          }
        },
        maintainAspectRatio: false,
        responsive: true,
        scales: {
          y: {
            beginAtZero: true,
            ticks: {
              stepSize: 10000000,
              callback: function(value) {
                if (Number(value) >= 1000000000) return Number(value) / 1000000000 + 'B';
                if (Number(value) >= 1000000) return Number(value) / 1000000 + 'M';
                if (Number(value) >= 1000) return Number(value) / 1000 + 'K';
                return Number(value);
              }
            }
          }
        }
      }
    });
  }, 400);

  export function updateData() {
    if (chart) {
      chart.data.labels = (chartRevenueReports || []).map((/** @type {ChartRevenueReport} */ i) => {
        const dt = /** @type {Date} */ (new Date(i.date));
        return dt.toLocaleDateString('en-US', {
          month: 'short',
          day: '2-digit'
        });
      });
      chart.data.datasets[0].data = (chartRevenueReports || []).map((/** @type {ChartRevenueReport} */ i) => Number(i.revenueIDR));
      chart.data.datasets[0].label = 'Revenue';
      chart.options.scales.y = {
        ticks: {
          stepSize: 10000000,
          callback: function(value) {
            if (Number(value) >= 1000000000) return Number(value) / 1000000000 + 'B';
            if (Number(value) >= 1000000) return Number(value) / 1000000 + 'M';
            if (Number(value) >= 1000) return Number(value) / 1000 + 'K';
            return Number(value);
          }
        }
      };
      chart.update();
    }
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

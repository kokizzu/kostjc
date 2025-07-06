<script>
  /**
   * @typedef {Object} Revenue
   * @property {string|number} revenue
   * @property {number} bookingTotal
   * @property {string} salesDate
   */

  import Chart from 'chart.js/auto';  

  let chart = /** @type {import('chart.js').Chart} */ (null);

  let revenues = /** @type {Revenue[]} */ ([
    {
      revenue: 560,
      bookingTotal: 4,
      salesDate: '2025-07-07'
    },
    {
      revenue: 890,
      bookingTotal: 8,
      salesDate: '2025-07-06'
    },
  ]);

  setTimeout(() => {
    const ElmChart = /** @type {HTMLCanvasElement} */ (document.getElementById('chart'));
    chart = new Chart(ElmChart, {
      type: 'line',
      data: {
        labels: (revenues || []).map((/** @type {Revenue} */ i) => {
          const dt = /** @type {Date} */ (new Date(i.salesDate));
          return dt.toLocaleDateString('en-US', {
            month: 'short',
            day: '2-digit'
          });
        }),
        datasets: [{
          label: 'Revenue',
          data: (revenues || []).map((/** @type {Revenue} */ i) => Number(i.revenue)),
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
</script>

<div class="chart">
  <canvas id="chart"></canvas>
</div>

<style>
  .chart {
    height: 92%;
    width: 100%;
    display: flex;
    width: 100%;
    padding: 0 16px 45px 16px;
  }

  .chart canvas {
    width: 100% !important;
    height: 100% !important;
  }
</style>

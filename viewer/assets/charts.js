google.charts.load('current', {'packages':['corechart']});
google.charts.setOnLoadCallback(drawChart);

function drawChart() {

    var xmlHttp = new XMLHttpRequest();
    xmlHttp.open( "GET", "http://localhost:8080/candlesticks", false );
    xmlHttp.send( null );

    var data = JSON.parse(xmlHttp.responseText);

    var output = data.map(function(obj) {
        return [obj.Day, obj.Min, obj.Open, obj.Close, obj.Max]
      });

    var data = google.visualization.arrayToDataTable(output, true);

    var options = {
      legend:'Bitcointrade Charts', 
      candlestick: {
          fallingColor: {
              fill: "#ff0000",
              stroke: "#ff0000"
          },
          risingColor: {
              fill: "#00cc00",
              stroke: "#00cc00"
          }
        },
        vAxis: {
            logScale: true,
        }

    };

    var chart = new google.visualization.CandlestickChart(document.getElementById('chart_div'));

    chart.draw(data, options);
  }
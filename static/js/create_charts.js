/**
 * Created by yannis on 3/6/2017.
 */
/**
 * Request data from the server, add it to the graph and set a timeout
 * to request again
 */

function createChart(device, container) {
    var chart; // global
    newChart(container);

    console.log(device);


    function requestData() {
        $.getJSON("/api/measurements/get/" + device, function (measurements) {
            // console.log(measurements);
            var series = chart.series[0];
            // console.log(typeof measurements);
            $.each(measurements, function (indexq, itemq) {
                // console.log(typeof itemq.Temperature);
                // console.log(itemq.Temperature);
                // console.log(typeof itemq.Timestamp);
                var shift = series.data.length > 200;// shift if the series is
                // longer than 20
                // console.log(new Date(itemq.Timestamp).getTime());

                chart.series[0].addPoint([new Date(itemq.Timestamp).getTime(), itemq.Temperature], true, shift);
            });


            // add the point


            // call it again after one second
            // setTimeout(requestData, 1000);
        });
    }


    function liveData() {

    }


    function newChart(container) {
        chart = new Highcharts.Chart({
            chart: {
                renderTo: container,
                defaultSeriesType: 'spline',
                events: {
                    load: requestData
                },
                zoomType: 'x'

            },
            title: {
                text: 'Temperature data'
            },
            xAxis: {
                type: 'datetime',
                tickPixelInterval: 150,
                labels: {
                    formatter: function () {
                        return Highcharts.dateFormat('%a %d %b %H:%M:%S', new Date(this.value));
                    }
                }
            },
            yAxis: {
                minPadding: 0.2,
                maxPadding: 0.2,
                title: {
                    text: 'Value',
                    margin: 80
                }
            },
            tooltip: {
                formatter: function () {
                    return '<b>' + this.series.name + '</b><br/>' +
                        Highcharts.dateFormat('%Y-%m-%d %H:%M:%S', this.x) + '<br/>' +
                        Highcharts.numberFormat(this.y, 2);
                }
            },
            scrollbar: {
                enabled: true
            },
            series: [{
                name: 'Temperature',
                data: []
            }]
        });
    }
}

// }
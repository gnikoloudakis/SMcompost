/**
 * Created by yannis on 3/6/2017.
 */
/**
 * Request data from the server, add it to the graph and set a timeout
 * to request again
 */

var wdata;
function createChart(device, container) {

    var chart; // global

    newChart(container);
    // requestData(device);
    // console.log(device);


    function requestData() {
        $.getJSON("/api/measurements/get/" + device, function (measurements) {
            $.each(measurements, function (indexq, itemq) {
                // var shift = series.data.length > 80;// shift if the series is
                chart.series[0].addPoint([new Date(itemq.Timestamp).getTime(), itemq.Temperature], false, false);
            });
            chart.redraw();
        });
    }


    function newChart() {
        Highcharts.setOptions({  // This is for all plots, change Date axis to local timezone
            global: {
                // useUTC : false,
                timezoneOffset: -4 * 60
            }
        });
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
            series: [{
                name: 'Temperature',
                data: []
            }]
        });

    }
}


function createChartByDate(device, container, startDate, stopDate) {
    var chart; // global
    newChart(container);

    // console.log(device);


    function requestData() {
        $.getJSON("/api/measurements/get/" + device + "/" + startDate + "/" + stopDate, function (measurements) {
            // console.log(measurements);
            var series = chart.series[0];
            // console.log(typeof measurements);
            $.each(measurements, function (indexq, itemq) {
                // var shift = series.data.length > 80;// shift if the series is
                chart.series[0].addPoint([new Date(itemq.Timestamp).getTime(), itemq.Temperature], false, false);
            });
            chart.redraw();
        });
    }


    function newChart(container) {
        Highcharts.setOptions({                                            // This is for all plots, change Date axis to local timezone
            global: {
                // useUTC : false,
                timezoneOffset: -4 * 60
            }
        });
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
            series: [{
                name: 'Temperature',
                data: []
            }]
        });
    }
}

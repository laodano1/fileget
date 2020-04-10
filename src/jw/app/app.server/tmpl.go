package main

var hptemplate = `
<!doctype html>
<html lang="en">
  <head>
    <!-- Required meta tags -->
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
	<link rel="icon" href="https://cn.bing.com/sa/simg/bing_p_rr_teal_min.ico" >

    <!-- Bootstrap CSS -->
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/css/bootstrap.min.css" integrity="sha384-Vkoo8x4CGsO3+Hhxv8T/Q5PaXtkKtu6ug5TOeNV6gBiFeWPGFN9MuhOf23Q9Ifjh" crossorigin="anonymous">

    <title>{{ .title }}</title>
  </head>
  <body>
    <h3 style="margin: 2px;">This is Sample Tool!</h3>
	<div style="left:5%;margin: 2px;"><b>Total client number: {{ .clientNum }}</b></div>
	<div class="container-fluid container-lg" style="left:5%;margin: 2px;">
		{{range $idx, $status := .allStatus}}
		<div class="card border-dark mb-3" style="max-width: 18rem;">
		  <div class="card-header">Client: {{ $status.cid }}</div>
		  <div class="card-body text-dark">
			<h5 class="card-title"></h5>
			<p class="card-text"> Phase: {{ $status.phase }} </p>
		  </div>
		</div>
		{{end}}
	</div>
		
 
    <!-- Optional JavaScript -->
    <!-- jQuery first, then Popper.js, then Bootstrap JS -->
    <script src="https://code.jquery.com/jquery-3.4.1.slim.min.js" integrity="sha384-J6qa4849blE2+poT4WnyKhv5vZF5SrPo0iEjwBvKU7imGFAV0wwj1yYfoRSJoZ+n" crossorigin="anonymous"></script>
    <script src="https://cdn.jsdelivr.net/npm/popper.js@1.16.0/dist/umd/popper.min.js" integrity="sha384-Q6E9RHvbIyZFJoft+2mJbHaEWldlvI9IOYy5n3zV9zzTtmI3UksdQRVvoxMfooAo" crossorigin="anonymous"></script>
    <script src="https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/js/bootstrap.min.js" integrity="sha384-wfSDF2E50Y2D1uUdj0O3uMBJnjuUD4Ih7YwaYd1iqfktj0Uod8GCExl3Og8ifwB6" crossorigin="anonymous"></script>
	<script>
		
		setInterval(function() {
			var wd = $( "progress-bar" ).width();
			console.log("progress width: ", wd)
			$( "progress-bar" ).width(wd + 2);
		}, 2000)
    </script>
  </body>
</html>

`


//<!--div class="row">
//<div class="col-sm-6">
//<ul class="list-group">
//<li class="list-group-item d-flex justify-content-between align-items-center">
//<b>Total client number:</b>
//<span class="badge badge-primary badge-pill">10</span>
//</li>
//
//<li class="list-group-item d-flex justify-content-between align-items-center" >
//<font size="2">&nbsp;&nbsp;&nbsp;Client id: {{ $status.cid }}</font>
//<div class="progress" >
//<div id="pb-{{ $idx }}" class="progress-bar progress-bar-striped progress-bar-animated" role="progressbar" style="width:10px" aria-valuenow="10" aria-valuemin="0" aria-valuemax="5"></div>
//</div>
//</li>
//<li class="list-group-item d-flex justify-content-between align-items-center">
//<font size="2">&nbsp;&nbsp;&nbsp;Client {{ $status.cid }} phase:</font>
//<span class="badge badge-primary badge-pill"> {{ $status.phase }} </span>
//<!--div class="progress" >
//<div id="pb-{{ $idx }}" class="progress-bar progress-bar-striped progress-bar-animated" role="progressbar" style="width:10px" aria-valuenow="10" aria-valuemin="0" aria-valuemax="5"></div>
//</div-->
//</li>
//{{end}}
//</ul>
//</div>
//<div class="col-sm-5">
//<textarea class="form-control" id="validationTextarea" placeholder="" rows="3" disabled>
//{{range $status := .allStatus}}
//{{ $status.log }}
//{{end}}
//</textarea>
//</div>
//<div class="col-sm-1">
//</div-->

<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="utf-8">
		<link href="bootstrap.min.css" rel="stylesheet">
		<title>Simple Test</title>
	</head>
	<!-- !! navbar-expand-md !! -->
	<body>
		<nav class="navbar bg-primary navbar-dark navbar-expand-md fixed-top">
			<a href="#" class="navbar-brand">Simple Test</a>
			<button type="button" class="navbar-toggler navbar-toggler-right" data-toggle="collapse" data-target="#navbar">
				<span class="navbar-toggler-icon"></span>
			</button>
			<div class="collapse navbar-collapse" id="navbar">
				<ul class="navbar-nav ml-auto">
					<li class="nav-item">
						<a class="nav-link" href="#Intro">Intro</a>
					</li>
					<li class="nav-item">
						<a class="nav-link" href="#Tests">Tests</a>
					</li>
					<li class="nav-item">
						<a class="nav-link" href="#Records">Records</a>
					</li>
					<li class="nav-item">
						<a class="nav-link" href="#Porfile">Profile</a>
					</li>
					<li class="nav-item">
						<a class="nav-link" href="#Login" data-toggle="modal" data-target="#modalLogin">Login</a>
					</li>
				</ul>
			</div>
		</nav>
		<div class="modal fade" id="modalLogin" tabindex="-1" role="dialog">
			<div class="modal-dialog modal-md" role="document">
				<div class="modal-content">
					<div class="modal-header">
						<h5 class="modal-title" id="loginTitle">Login</h5>
						<button type="button" class="close" data-dismiss="modal" aria-label="Close">
              <span aria-hidden="true">&times;</span>
            </button>
					</div>
					<div class="modal-body">
						<form action="/login/" method="POST">
							<div class="form-group">
		            <label for="input-username" class="sr-only">Username:</label>
		            <input type="text" class="form-control" id="input-username" name="username">
		          </div>
		          <div class="form-group">
		            <label for="input-password" class="sr-only">Password:</label>
		            <input type="password" class="form-control" id="input-password">
		          </div>
		          <div class="form-check">
		            <label class="form-check-label">
		              <input type="checkbox" class="form-check-input" id="input-keepLogin">
		              Keep Login
		              <small class="form-text">Reminder: Don't use Keep Login on a untrusted computer.</small>
		            </label>
		          </div>
		          <div>
		            <button type="submit" class="btn btn-dark">Login</button>
		          </div>
						</form>
					</div>
				</div>
			</div>
		</div>

		<section class="jumbotron jumbotron-fluid">
			<div class="container">
				<h1 class="diplay-1">Simple Test</h1>
			</div>
		</section> 

		<script src="jquery-3.2.1.min.js"></script>
		<script src="popper.js"></script>
		<script src="bootstrap.min.js"></script>
	</body>
</html>
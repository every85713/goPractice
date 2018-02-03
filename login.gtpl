<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="utf-8">
		<link href="css/bootstrap.min.css" rel="stylesheet">
		<link href="css/style.css" rel="stylesheet">
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
		<div id="login-form">
			<h2 class="display-5">Login</h2>
			<section id="panel">
				<form action="/login/" method="POST">
					<div class="form-group">
            <label for="input-username" class="sr-only">Username:</label>
            <input type="text" class="form-control" id="input-username" name="username" placeholder="Username">
          </div>
          <div class="form-group">
            <label for="input-password" class="sr-only">Password:</label>
            <input type="password" class="form-control" name="password" id="input-password" placeholder="Password">
          </div>
          <div class="form-check">
            <label class="form-check-label">
              <input type="checkbox" class="form-check-input" id="input-keepLogin">
              Keep Login
              <small class="form-text">Reminder: Don't use Keep Login on a untrusted computer.</small>
            </label>
          </div>
          <br>
          <div>
            <button type="submit" class="btn btn-dark" value="login">Login</button>
          </div>
				</form>
				<br>
				<div style="display: block; text-align: right;">
					<a href="#">Forget Your Account?</a>
				</div>
			</section>
			<section id="panel">
				<form action="/login/" method="POST">
					<div class="form-group">
            <label for="input-username" class="sr-only">Username:</label>
            <input type="text" class="form-control" id="input-username" name="username" placeholder="Keeped Area">
          </div>
          <div class="form-group">
            <label for="input-password" class="sr-only">Password:</label>
            <input type="password" class="form-control" id="input-password" placeholder="For FB/Google+/Twitter Login">
          </div>
          <div class="form-check">
            <label class="form-check-label">
              <input type="checkbox" class="form-check-input" id="input-keepLogin">
              Keep Login
              <small class="form-text">Reminder: Don't use Keep Login on a untrusted computer.</small>
            </label>
          </div>
          <br>
          <div>
            <button type="submit" class="btn btn-dark">Login</button>
          </div>
				</form>
			</section>
		</div>


		<script src="js/jquery-3.2.1.min.js"></script>
		<script src="js/popper.js"></script>
		<script src="js/bootstrap.min.js"></script>
	</body>
</html>
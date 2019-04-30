{{define "navbar"}}
<a class="navbar-brand" href="/">My Blog</a>
<div>
  <ul class="nav navbar-nav">
    <li {{if .IsHome}}class="active"{{end}}><a href="/">HOME</a></li>
    <li {{if .IsCategory}}class="active"{{end}}><a href="/category">Category</a></li>
    <li {{if .IsTopic}}class="active"{{end}}><a href="/topic">Topic</a></li>
  </ul>
</div>

<div class="pull-right">
	<ul class="nav navbar-nav">
		{{if .IsLogined}}
		<li><a href="/login?exitLogin=true">Sign out</a></li>
		{{else}}
		<li><a href="/login">Sign in or Sign up</a></li>
		{{end}}
	</ul>
</div>
{{end}}
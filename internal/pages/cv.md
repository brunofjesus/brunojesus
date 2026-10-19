<style>
/* Timeline */
.timeline {
	margin-top: 20px;
	padding-left: 24px;
	border-left: 2px solid #cda869;
}

.timeline .item {
	position: relative;
	display: grid;
	grid-template-columns: 11em 1fr;
	grid-template-areas:
		"meta place"
		"meta desc";
	column-gap: 24px;
	padding-bottom: 24px;
}

.timeline .item:last-child {
	padding-bottom: 0;
}

.timeline .item:before {
	content: "";
	position: absolute;
	top: 4px;
	left: -31px;
	width: 10px;
	height: 10px;
	border-radius: 50%;
	background-color: #1e1e1e;
	border: 2px solid #cda869;
}

.timeline .work-place {
	grid-area: place;
	font-weight: bold;
}

.timeline .job-meta {
	grid-area: meta;
	text-align: right;
}

.timeline .job-meta .time {
	color: #7e7e7e;
	font-size: 11pt;
}

.timeline .job-desc {
	grid-area: desc;
	color: #cfcfcf;
}

.timeline .job-desc p {
	margin: 4px 0 0;
}

@media only screen and (max-width: 600px) {
	.timeline .item {
		grid-template-columns: 1fr;
		grid-template-areas:
			"meta"
			"place"
			"desc";
	}

	.timeline .job-meta {
		text-align: left;
		margin-bottom: 4px;
	}
}

/* Skills */
.skills {
	display: grid;
	grid-template-columns: repeat(3, 1fr);
	gap: 24px;
	margin-top: 20px;
}

.skills .skill {
	text-align: center;
}

.skills .skill .language {
	font-weight: bold;
	margin-bottom: 4px;
}

.skills .skill p {
	color: #cfcfcf;
	font-size: 11pt;
}
</style>

# Curriculum Vitae
<div class="button-container">
  <a href="assets/cv.pdf" download class="cv">
    <img src="assets/svg/file.svg" alt="LinkedIn"
      style="width: 20px; height: 20px; filter: brightness(0) invert(1);">
    Download CV
  </a>

  <a href="https://github.com/brunofjesus" class="github" target="_blank">
    <img src="assets/svg/github.svg" alt="LinkedIn"
      style="width: 20px; height: 20px; filter: brightness(0) invert(1);">
    GitHub
  </a>

  <a href="https://www.linkedin.com/in/bruno-jesus-0ab943a9" class="linkedin" target="_blank">
    <img src="assets/svg/linkedin.svg" alt="LinkedIn"
      style="width: 20px; height: 20px; filter: brightness(0) invert(1);">
    LinkedIn
  </a>
</div>


I'm Bruno Jesus, a software developer from Portugal. I've been writing software
professionally since 2012, and these days I mostly write Go and spend my time
around Kubernetes, platform tooling, and developer experience.

Here's a rough timeline of the places I've worked and the things I've done.

## Work experience

<div class="timeline">
  <div class="item">
    <div class="work-place">commercetools</div>
    <div class="job-meta">
      <div class="title">Senior Software Engineer</div>
      <div class="time">2022-01 / Present</div>
    </div>
    <div class="job-desc">
      <p>Development of <a href="https://commercetools.com/features/frontend">commercetools Frontend</a>.
      Led the transition to multi-tenancy with Kubernetes, and the design and implementation of telemetry and observability.</p>
    </div>
  </div>

  <div class="item">
    <div class="work-place">Escola Superior de Gestão e Tecnologia de Santarém</div>
    <div class="job-meta">
      <div class="title">Invited Assistant</div>
      <div class="time">2021-10 / 2024-07</div>
      <div class="time">2026-03 / Present</div>
    </div>
    <div class="job-desc">
      <p><b>Taught Courses:</b> Relational Databases, Software Engineering, Systems Integration, Digital Systems, and Operating Systems.</p>
    </div>
  </div>

  <div class="item">
    <div class="work-place">Mercedes-Benz.io @ Cloudmobility</div>
    <div class="job-meta">
      <div class="title">Senior Software Engineer</div>
      <div class="time">2021-08 / 2022-02</div>
    </div>
    <div class="job-desc">
      <p>Development of cNS (Cloudmobility Namespaces), a private cloud for Daimler.</p>
    </div>
  </div>

  <div class="item">
    <div class="work-place">Volkswagen Digital Solutions @ MAN Digital Hub</div>
    <div class="job-meta">
      <div class="title">Software Engineer</div>
      <div class="time">2020-03 / 2021-08</div>
    </div>
    <div class="job-desc">
      <p>Development of a Manufacturing Execution System for MAN Truck &amp; Bus.</p>
    </div>
  </div>

  <div class="item">
    <div class="work-place">Alter Solutions @ AXA Belgium</div>
    <div class="job-meta">
      <div class="title">Software Engineer</div>
      <div class="time">2019-06 / 2020-03</div>
    </div>
    <div class="job-desc">
      <p>Development of <a href="https://www.axa.be/myaxaaccidentatwork">MyAXA Accident@Work</a>.</p>
    </div>
  </div>

  <div class="item">
    <div class="work-place">Maxdata</div>
    <div class="job-meta">
      <div class="title">Software Engineer</div>
      <div class="time">2017-06 / 2019-06</div>
    </div>
    <div class="job-desc">
      <p>Development of <a href="https://www.maxdata.pt/">Clinidata</a> healthcare software.</p>
    </div>
  </div>

  <div class="item">
    <div class="work-place">Redislogar</div>
    <div class="job-meta">
      <div class="title">Freelancer Developer</div>
      <div class="time">2016-02 / 2016-04</div>
    </div>
    <div class="job-desc">
      <p>Configuration of a custom Linux distribution, system security, and development of an update system for their HomePaq machines.</p>
    </div>
  </div>

  <div class="item">
    <div class="work-place">ConnectedTI</div>
    <div class="job-meta">
      <div class="title">Intern</div>
      <div class="time">2014-06 / 2014-08</div>
    </div>
    <div class="job-desc">
      <p>Internship for my specialization course; development of a CRM for a call center.</p>
    </div>
  </div>

  <div class="item">
    <div class="work-place">Wolfd</div>
    <div class="job-meta">
      <div class="title">Junior Developer</div>
      <div class="time">2012-06 / 2013-09</div>
    </div>
    <div class="job-desc">
      <p>Most of my time was spent developing Rediskey, my first "big" project.</p>
    </div>
  </div>
</div>

## Skills

<div class="skills">
  <div class="skill">
    <div class="language">Go ❤️</div>
    <p>Fell in love with its ease of use, performance, and did I say it compiles to a single binary?</p>
  </div>
  <div class="skill">
    <div class="language">Java 👌</div>
    <p>Was part of my day to day life from 2017 to 2022. After that I mainly used it for teaching.</p>
  </div>
  <div class="skill">
    <div class="language">PHP 🤷</div>
    <p>Some systems at my current job use it, that means I have to use it as well.</p>
  </div>
  <div class="skill">
    <div class="language">TypeScript 😐</div>
    <p>Makes JavaScript bearable and gets the job done.</p>
  </div>
  <div class="skill">
    <div class="language">Python 🤓</div>
    <p>I don't really use it much nowadays.</p>
  </div>
  <div class="skill">
    <div class="language">Kubernetes 😇</div>
    <p>I guess I like it because I don't really have to maintain it!</p>
  </div>
</div>

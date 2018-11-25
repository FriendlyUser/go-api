<template>
  <div class="hello">
    <h1>Job Info</h1>
    <ul id="example-1">
	  <li v-for="job in jobs">
		{{ job.id }}
	  </li>
	</ul>
    <table class="table table-bordered table-hover table-condensed">
<thead><tr><th title="Field #1">id</th>
<th title="Field #2">numjobs</th>
<th title="Field #3">avgkeywords</th>
<th title="Field #4">avgskills</th>
<th title="Field #5">city</th>
<th title="Field #6">searchterm</th>
<th title="Field #7">searchtime</th>
</tr></thead>
<tbody><tr>
<td align="right">14</td>
<td align="right">409</td>
<td align="right">0.91</td>
<td align="right">1.58</td>
<td>Vancouver</td>
<td>Machine Learning</td>
<td>2018-20-09</td>
</tr>
<tr>
<td align="right">15</td>
<td align="right">181</td>
<td align="right">0.18</td>
<td align="right">2.3</td>
<td>Victoria</td>
<td>Developer</td>
<td>2018-20-09</td>
</tr>
<tr>
<td align="right">16</td>
<td align="right">95</td>
<td align="right">1.77</td>
<td align="right">1.65</td>
<td>Vancouver</td>
<td>Blockchain</td>
<td>2018-20-09</td>
</tr>
<tr>
<td align="right">17</td>
<td align="right">227</td>
<td align="right">0.29</td>
<td align="right">3.03</td>
<td>Vancouver</td>
<td>Junior Developer</td>
<td>2018-20-09</td>
</tr>
<tr>
<td align="right">18</td>
<td align="right">403</td>
<td align="right">0.95</td>
<td align="right">1.59</td>
<td>Vancouver</td>
<td>Machine Learning</td>
<td>2018-23-09</td>
</tr>
<tr>
<td align="right">6</td>
<td align="right">446</td>
<td align="right">0.91</td>
<td align="right">1.63</td>
<td>Vancouver</td>
<td>Machine Learning</td>
<td>2018-16-09</td>
</tr>
<tr>
<td align="right">7</td>
<td align="right">180</td>
<td align="right">0.25</td>
<td align="right">2.42</td>
<td>Victoria</td>
<td>Developer</td>
<td>2018-16-09</td>
</tr>
<tr>
<td align="right">8</td>
<td align="right">90</td>
<td align="right">1.76</td>
<td align="right">1.5</td>
<td>Vancouver</td>
<td>Blockchain</td>
<td>2018-16-09</td>
</tr>
<tr>
<td align="right">9</td>
<td align="right">214</td>
<td align="right">0.23</td>
<td align="right">2.81</td>
<td>Vancouver</td>
<td>Junior Developer</td>
<td>2018-16-09</td>
</tr>
<tr>
<td align="right">10</td>
<td align="right">446</td>
<td align="right">0.9</td>
<td align="right">1.54</td>
<td>Vancouver</td>
<td>Machine Learning</td>
<td>2018-16-09</td>
</tr>
<tr>
<td align="right">11</td>
<td align="right">180</td>
<td align="right">0.24</td>
<td align="right">2.39</td>
<td>Victoria</td>
<td>Developer</td>
<td>2018-16-09</td>
</tr>
<tr>
<td align="right">12</td>
<td align="right">90</td>
<td align="right">1.76</td>
<td align="right">1.5</td>
<td>Vancouver</td>
<td>Blockchain</td>
<td>2018-16-09</td>
</tr>
<tr>
<td align="right">13</td>
<td align="right">215</td>
<td align="right">0.26</td>
<td align="right">2.84</td>
<td>Vancouver</td>
<td>Junior Developer</td>
<td>2018-16-09</td>
</tr>
<tr>
<td align="right">19</td>
<td align="right">177</td>
<td align="right">0.17</td>
<td align="right">2.22</td>
<td>Victoria</td>
<td>Developer</td>
<td>2018-23-09</td>
</tr>
<tr>
<td align="right">20</td>
<td align="right">89</td>
<td align="right">1.79</td>
<td align="right">1.58</td>
<td>Vancouver</td>
<td>Blockchain</td>
<td>2018-23-09</td>
</tr>
<tr>
<td align="right">21</td>
<td align="right">208</td>
<td align="right">0.3</td>
<td align="right">3.1</td>
<td>Vancouver</td>
<td>Junior Developer</td>
<td>2018-23-09</td>
</tr>
<tr>
<td align="right">22</td>
<td align="right">410</td>
<td align="right">0.96</td>
<td align="right">1.57</td>
<td>Vancouver</td>
<td>Machine Learning</td>
<td>2018-27-09</td>
</tr>
<tr>
<td align="right">23</td>
<td align="right">178</td>
<td align="right">0.2</td>
<td align="right">2.28</td>
<td>Victoria</td>
<td>Developer</td>
<td>2018-27-09</td>
</tr>
<tr>
<td align="right">24</td>
<td align="right">89</td>
<td align="right">1.89</td>
<td align="right">1.64</td>
<td>Vancouver</td>
<td>Blockchain</td>
<td>2018-27-09</td>
</tr>
<tr>
<td align="right">25</td>
<td align="right">220</td>
<td align="right">0.29</td>
<td align="right">2.87</td>
<td>Vancouver</td>
<td>Junior Developer</td>
<td>2018-27-09</td>
</tr>
<tr>
<td align="right">26</td>
<td align="right">417</td>
<td align="right">0.92</td>
<td align="right">1.55</td>
<td>Vancouver</td>
<td>Machine Learning</td>
<td>2018-30-09</td>
</tr>
<tr>
<td align="right">27</td>
<td align="right">189</td>
<td align="right">0.21</td>
<td align="right">2.06</td>
<td>Victoria</td>
<td>Developer</td>
<td>2018-30-09</td>
</tr>
<tr>
<td align="right">28</td>
<td align="right">90</td>
<td align="right">1.93</td>
<td align="right">1.86</td>
<td>Vancouver</td>
<td>Blockchain</td>
<td>2018-30-09</td>
</tr>
<tr>
<td align="right">29</td>
<td align="right">223</td>
<td align="right">0.3</td>
<td align="right">2.74</td>
<td>Vancouver</td>
<td>Junior Developer</td>
<td>2018-30-09</td>
</tr>
<tr>
<td align="right">30</td>
<td align="right">336</td>
<td align="right">0.86</td>
<td align="right">1.43</td>
<td>Vancouver</td>
<td>Machine Learning</td>
<td>2018-04-10</td>
</tr>
<tr>
<td align="right">31</td>
<td align="right">180</td>
<td align="right">0.24</td>
<td align="right">2.03</td>
<td>Victoria</td>
<td>Developer</td>
<td>2018-04-10</td>
</tr>
<tr>
<td align="right">32</td>
<td align="right">81</td>
<td align="right">1.99</td>
<td align="right">1.95</td>
<td>Vancouver</td>
<td>Blockchain</td>
<td>2018-04-10</td>
</tr>
<tr>
<td align="right">33</td>
<td align="right">207</td>
<td align="right">0.27</td>
<td align="right">2.7</td>
<td>Vancouver</td>
<td>Junior Developer</td>
<td>2018-04-10</td>
</tr>
<tr>
<td align="right">34</td>
<td align="right">16</td>
<td align="right">0.1</td>
<td align="right">0.7</td>
<td>Victoria</td>
<td>Machine Learning</td>
<td>2018-07-10</td>
</tr>
<tr>
<td align="right">35</td>
<td align="right">426</td>
<td align="right">0.91</td>
<td align="right">1.53</td>
<td>Vancouver</td>
<td>Machine Learning</td>
<td>2018-07-10</td>
</tr>
<tr>
<td align="right">36</td>
<td align="right">192</td>
<td align="right">0.21</td>
<td align="right">2.09</td>
<td>Victoria</td>
<td>Developer</td>
<td>2018-07-10</td>
</tr>
<tr>
<td align="right">37</td>
<td align="right">93</td>
<td align="right">1.94</td>
<td align="right">1.91</td>
<td>Vancouver</td>
<td>Blockchain</td>
<td>2018-07-10</td>
</tr>
<tr>
<td align="right">38</td>
<td align="right">207</td>
<td align="right">0.25</td>
<td align="right">2.69</td>
<td>Vancouver</td>
<td>Junior Developer</td>
<td>2018-07-10</td>
</tr>
<tr>
<td align="right">39</td>
<td align="right">418</td>
<td align="right">0.88</td>
<td align="right">1.54</td>
<td>Vancouver</td>
<td>Machine Learning</td>
<td>2018-11-10</td>
</tr>
<tr>
<td align="right">40</td>
<td align="right">186</td>
<td align="right">0.2</td>
<td align="right">1.93</td>
<td>Victoria</td>
<td>Developer</td>
<td>2018-11-10</td>
</tr>
<tr>
<td align="right">41</td>
<td align="right">99</td>
<td align="right">1.86</td>
<td align="right">1.82</td>
<td>Vancouver</td>
<td>Blockchain</td>
<td>2018-11-10</td>
</tr>
<tr>
<td align="right">42</td>
<td align="right">194</td>
<td align="right">0.24</td>
<td align="right">2.81</td>
<td>Vancouver</td>
<td>Junior Developer</td>
<td>2018-11-10</td>
</tr>
<tr>
<td align="right">43</td>
<td align="right">335</td>
<td align="right">0.9</td>
<td align="right">1.54</td>
<td>Vancouver</td>
<td>Machine Learning</td>
<td>2018-14-10</td>
</tr>
<tr>
<td align="right">44</td>
<td align="right">178</td>
<td align="right">0.2</td>
<td align="right">1.94</td>
<td>Victoria</td>
<td>Developer</td>
<td>2018-14-10</td>
</tr>
<tr>
<td align="right">45</td>
<td align="right">94</td>
<td align="right">1.83</td>
<td align="right">1.56</td>
<td>Vancouver</td>
<td>Blockchain</td>
<td>2018-14-10</td>
</tr>
<tr>
<td align="right">46</td>
<td align="right">189</td>
<td align="right">0.26</td>
<td align="right">2.63</td>
<td>Vancouver</td>
<td>Junior Developer</td>
<td>2018-14-10</td>
</tr>
<tr>
<td align="right">47</td>
<td align="right">410</td>
<td align="right">0.85</td>
<td align="right">1.5</td>
<td>Vancouver</td>
<td>Machine Learning</td>
<td>2018-18-10</td>
</tr>
<tr>
<td align="right">48</td>
<td align="right">181</td>
<td align="right">0.17</td>
<td align="right">2.06</td>
<td>Victoria</td>
<td>Developer</td>
<td>2018-18-10</td>
</tr>
<tr>
<td align="right">49</td>
<td align="right">91</td>
<td align="right">1.75</td>
<td align="right">1.71</td>
<td>Vancouver</td>
<td>Blockchain</td>
<td>2018-18-10</td>
</tr>
<tr>
<td align="right">50</td>
<td align="right">190</td>
<td align="right">0.24</td>
<td align="right">2.57</td>
<td>Vancouver</td>
<td>Junior Developer</td>
<td>2018-18-10</td>
</tr>
<tr>
<td align="right">51</td>
<td align="right">336</td>
<td align="right">0.98</td>
<td align="right">1.78</td>
<td>Vancouver</td>
<td>Machine Learning</td>
<td>2018-18-10</td>
</tr>
<tr>
<td align="right">52</td>
<td align="right">183</td>
<td align="right">0.2</td>
<td align="right">2.04</td>
<td>Victoria</td>
<td>Developer</td>
<td>2018-18-10</td>
</tr>
<tr>
<td align="right">53</td>
<td align="right">86</td>
<td align="right">1.85</td>
<td align="right">1.64</td>
<td>Vancouver</td>
<td>Blockchain</td>
<td>2018-18-10</td>
</tr>
<tr>
<td align="right">54</td>
<td align="right">191</td>
<td align="right">0.19</td>
<td align="right">2.47</td>
<td>Vancouver</td>
<td>Junior Developer</td>
<td>2018-18-10</td>
</tr>
<tr>
<td align="right">84</td>
<td align="right">404</td>
<td align="right">0.9</td>
<td align="right">1.84</td>
<td>Vancouver</td>
<td>Machine Learning</td>
<td>2018-20-10</td>
</tr>
<tr>
<td align="right">85</td>
<td align="right">177</td>
<td align="right">0.14</td>
<td align="right">2.05</td>
<td>Victoria</td>
<td>Developer</td>
<td>2018-20-10</td>
</tr>
<tr>
<td align="right">86</td>
<td align="right">88</td>
<td align="right">1.81</td>
<td align="right">1.83</td>
<td>Vancouver</td>
<td>Blockchain</td>
<td>2018-20-10</td>
</tr>
<tr>
<td align="right">87</td>
<td align="right">196</td>
<td align="right">0.28</td>
<td align="right">2.55</td>
<td>Vancouver</td>
<td>Junior Developer</td>
<td>2018-20-10</td>
</tr>
<tr>
<td align="right">88</td>
<td align="right">402</td>
<td align="right">0.88</td>
<td align="right">1.7</td>
<td>Vancouver</td>
<td>Machine Learning</td>
<td>2018-21-10</td>
</tr>
<tr>
<td align="right">89</td>
<td align="right">188</td>
<td align="right">0.19</td>
<td align="right">2.01</td>
<td>Victoria</td>
<td>Developer</td>
<td>2018-21-10</td>
</tr>
<tr>
<td align="right">90</td>
<td align="right">87</td>
<td align="right">1.75</td>
<td align="right">1.64</td>
<td>Vancouver</td>
<td>Blockchain</td>
<td>2018-21-10</td>
</tr>
<tr>
<td align="right">91</td>
<td align="right">193</td>
<td align="right">0.31</td>
<td align="right">2.8</td>
<td>Vancouver</td>
<td>Junior Developer</td>
<td>2018-21-10</td>
</tr>
<tr>
<td align="right">92</td>
<td align="right">375</td>
<td align="right">0.89</td>
<td align="right">1.76</td>
<td>Vancouver</td>
<td>Machine Learning</td>
<td>2018-25-10</td>
</tr>
<tr>
<td align="right">93</td>
<td align="right">187</td>
<td align="right">0.23</td>
<td align="right">1.98</td>
<td>Victoria</td>
<td>Developer</td>
<td>2018-25-10</td>
</tr>
<tr>
<td align="right">94</td>
<td align="right">87</td>
<td align="right">1.7</td>
<td align="right">1.72</td>
<td>Vancouver</td>
<td>Blockchain</td>
<td>2018-25-10</td>
</tr>
<tr>
<td align="right">95</td>
<td align="right">186</td>
<td align="right">0.27</td>
<td align="right">2.68</td>
<td>Vancouver</td>
<td>Junior Developer</td>
<td>2018-25-10</td>
</tr>
<tr>
<td align="right">96</td>
<td align="right">373</td>
<td align="right">0.89</td>
<td align="right">1.7</td>
<td>Vancouver</td>
<td>Machine Learning</td>
<td>2018-25-10</td>
</tr>
<tr>
<td align="right">97</td>
<td align="right">187</td>
<td align="right">0.22</td>
<td align="right">2.14</td>
<td>Victoria</td>
<td>Developer</td>
<td>2018-25-10</td>
</tr>
<tr>
<td align="right">98</td>
<td align="right">87</td>
<td align="right">1.68</td>
<td align="right">1.67</td>
<td>Vancouver</td>
<td>Blockchain</td>
<td>2018-25-10</td>
</tr>
<tr>
<td align="right">99</td>
<td align="right">189</td>
<td align="right">0.29</td>
<td align="right">2.63</td>
<td>Vancouver</td>
<td>Junior Developer</td>
<td>2018-25-10</td>
</tr>
<tr>
<td align="right">100</td>
<td align="right">355</td>
<td align="right">0.92</td>
<td align="right">1.7</td>
<td>Vancouver</td>
<td>Machine Learning</td>
<td>2018-01-11</td>
</tr>
<tr>
<td align="right">101</td>
<td align="right">355</td>
<td align="right">0.92</td>
<td align="right">1.73</td>
<td>Vancouver</td>
<td>Machine Learning</td>
<td>2018-01-11</td>
</tr>
<tr>
<td align="right">102</td>
<td align="right">93</td>
<td align="right">1.71</td>
<td align="right">1.65</td>
<td>Vancouver</td>
<td>Blockchain</td>
<td>2018-01-11</td>
</tr>
<tr>
<td align="right">103</td>
<td align="right">355</td>
<td align="right">0.93</td>
<td align="right">1.71</td>
<td>Vancouver</td>
<td>Machine Learning</td>
<td>2018-01-11</td>
</tr>
<tr>
<td align="right">104</td>
<td align="right">93</td>
<td align="right">1.72</td>
<td align="right">1.59</td>
<td>Vancouver</td>
<td>Blockchain</td>
<td>2018-01-11</td>
</tr>
<tr>
<td align="right">105</td>
<td align="right">193</td>
<td align="right">0.31</td>
<td align="right">2.74</td>
<td>Vancouver</td>
<td>Junior Developer</td>
<td>2018-01-11</td>
</tr>
<tr>
<td align="right">106</td>
<td align="right">187</td>
<td align="right">0.19</td>
<td align="right">1.95</td>
<td>Victoria</td>
<td>Developer</td>
<td>2018-01-11</td>
</tr>
<tr>
<td align="right">107</td>
<td align="right">347</td>
<td align="right">0.98</td>
<td align="right">1.83</td>
<td>Vancouver</td>
<td>Machine Learning</td>
<td>2018-04-11</td>
</tr>
<tr>
<td align="right">108</td>
<td align="right">87</td>
<td align="right">1.64</td>
<td align="right">1.62</td>
<td>Vancouver</td>
<td>Blockchain</td>
<td>2018-04-11</td>
</tr>
<tr>
<td align="right">109</td>
<td align="right">199</td>
<td align="right">0.29</td>
<td align="right">2.67</td>
<td>Vancouver</td>
<td>Junior Developer</td>
<td>2018-04-11</td>
</tr>
<tr>
<td align="right">110</td>
<td align="right">185</td>
<td align="right">0.18</td>
<td align="right">2.01</td>
<td>Victoria</td>
<td>Developer</td>
<td>2018-04-11</td>
</tr>
<tr>
<td align="right">111</td>
<td align="right">355</td>
<td align="right">0.96</td>
<td align="right">1.89</td>
<td>Vancouver</td>
<td>Machine Learning</td>
<td>2018-08-11</td>
</tr>
<tr>
<td align="right">112</td>
<td align="right">79</td>
<td align="right">1.58</td>
<td align="right">1.35</td>
<td>Vancouver</td>
<td>Blockchain</td>
<td>2018-08-11</td>
</tr>
<tr>
<td align="right">113</td>
<td align="right">199</td>
<td align="right">0.26</td>
<td align="right">2.7</td>
<td>Vancouver</td>
<td>Junior Developer</td>
<td>2018-08-11</td>
</tr>
<tr>
<td align="right">114</td>
<td align="right">172</td>
<td align="right">0.15</td>
<td align="right">2.12</td>
<td>Victoria</td>
<td>Developer</td>
<td>2018-08-11</td>
</tr>
<tr>
<td align="right">115</td>
<td align="right">353</td>
<td align="right">0.97</td>
<td align="right">1.89</td>
<td>Vancouver</td>
<td>Machine Learning</td>
<td>2018-11-11</td>
</tr>
<tr>
<td align="right">116</td>
<td align="right">65</td>
<td align="right">1.63</td>
<td align="right">1.57</td>
<td>Vancouver</td>
<td>Blockchain</td>
<td>2018-11-11</td>
</tr>
<tr>
<td align="right">117</td>
<td align="right">193</td>
<td align="right">0.27</td>
<td align="right">2.41</td>
<td>Vancouver</td>
<td>Junior Developer</td>
<td>2018-11-11</td>
</tr>
<tr>
<td align="right">118</td>
<td align="right">174</td>
<td align="right">0.15</td>
<td align="right">2.07</td>
<td>Victoria</td>
<td>Developer</td>
<td>2018-11-11</td>
</tr>
<tr>
<td align="right">119</td>
<td align="right">358</td>
<td align="right">0.94</td>
<td align="right">1.91</td>
<td>Vancouver</td>
<td>Machine Learning</td>
<td>2018-15-11</td>
</tr>
<tr>
<td align="right">120</td>
<td align="right">72</td>
<td align="right">1.58</td>
<td align="right">1.48</td>
<td>Vancouver</td>
<td>Blockchain</td>
<td>2018-15-11</td>
</tr>
<tr>
<td align="right">121</td>
<td align="right">189</td>
<td align="right">0.29</td>
<td align="right">2.2</td>
<td>Vancouver</td>
<td>Junior Developer</td>
<td>2018-15-11</td>
</tr>
<tr>
<td align="right">122</td>
<td align="right">179</td>
<td align="right">0.14</td>
<td align="right">1.98</td>
<td>Victoria</td>
<td>Developer</td>
<td>2018-15-11</td>
</tr>
<tr>
<td align="right">123</td>
<td align="right">361</td>
<td align="right">0.93</td>
<td align="right">1.88</td>
<td>Vancouver</td>
<td>Machine Learning</td>
<td>2018-18-11</td>
</tr>
<tr>
<td align="right">124</td>
<td align="right">79</td>
<td align="right">1.57</td>
<td align="right">1.66</td>
<td>Vancouver</td>
<td>Blockchain</td>
<td>2018-18-11</td>
</tr>
<tr>
<td align="right">125</td>
<td align="right">200</td>
<td align="right">0.26</td>
<td align="right">2.37</td>
<td>Vancouver</td>
<td>Junior Developer</td>
<td>2018-18-11</td>
</tr>
<tr>
<td align="right">126</td>
<td align="right">189</td>
<td align="right">0.16</td>
<td align="right">1.99</td>
<td>Victoria</td>
<td>Developer</td>
<td>2018-18-11</td>
</tr>
<tr>
<td align="right">127</td>
<td align="right">368</td>
<td align="right">0.97</td>
<td align="right">1.89</td>
<td>Vancouver</td>
<td>Machine Learning</td>
<td>2018-22-11</td>
</tr>
<tr>
<td align="right">128</td>
<td align="right">84</td>
<td align="right">1.41</td>
<td align="right">1.66</td>
<td>Vancouver</td>
<td>Blockchain</td>
<td>2018-22-11</td>
</tr>
<tr>
<td align="right">129</td>
<td align="right">208</td>
<td align="right">0.28</td>
<td align="right">2.43</td>
<td>Vancouver</td>
<td>Junior Developer</td>
<td>2018-22-11</td>
</tr>
<tr>
<td align="right">130</td>
<td align="right">175</td>
<td align="right">0.16</td>
<td align="right">2.12</td>
<td>Victoria</td>
<td>Developer</td>
<td>2018-22-11</td>
</tr>
<tr>
<td align="right">131</td>
<td align="right">357</td>
<td align="right">0.97</td>
<td align="right">1.92</td>
<td>Vancouver</td>
<td>Machine Learning</td>
<td>2018-25-11</td>
</tr>
<tr>
<td align="right">132</td>
<td align="right">86</td>
<td align="right">1.36</td>
<td align="right">1.55</td>
<td>Vancouver</td>
<td>Blockchain</td>
<td>2018-25-11</td>
</tr>
<tr>
<td align="right">133</td>
<td align="right">207</td>
<td align="right">0.28</td>
<td align="right">2.54</td>
<td>Vancouver</td>
<td>Junior Developer</td>
<td>2018-25-11</td>
</tr>
<tr>
<td align="right">134</td>
<td align="right">184</td>
<td align="right">0.17</td>
<td align="right">2.39</td>
<td>Victoria</td>
<td>Developer</td>
<td>2018-25-11</td>
</tr>
</tbody></table>
    <br>
    <hr>
    <br>
  </div>
</template>

<script>
import axios from 'axios';

const api = 'https://golang-job-api.herokuapp.com/api';
// import { VueGoodTable } from 'vue-good-table';

export default {
  name: 'hello',
  data: () => ({
    productName: null,
    productPrice: 0.0,
    jobs: [],
  }),
  // add to component
  // components: {
  //   VueGoodTable,
  // },

  methods: {
    /**
    async createProduct() {
      await axios.post(`${api}/jobs`, {
        name: this.productName,
        price: Number(this.productPrice),
      });

      // refresh the data
      this.retrieveProducts();
    },

    async deleteProduct(id) {
      // delete the product
      await axios.delete(`${api}/jobs/${id}`);

      // refresh the data
      const response = await axios.get(`${api}/jobs`);
      this.products = response.data;
    },
    */
    async retrieveProducts() {
      const response = await axios.get(`${api}/jobs`);
      console.log(response)
      this.jobs = response.data.sort((a, b) => a.id - b.id);
    },
  },

  async created() {
    console.log('wait')
    this.retrieveProducts();
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
}

ul {
  list-style-type: none;
  padding: 0;
  padding-left: 40%;
}

li {
  /*display: inline-block;*/
  margin: 0 10px;
  text-align: left;
}

a {
  color: #42b983;
}
</style>

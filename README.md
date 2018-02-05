<h1>A sample application using hyperledger fabric</h1>

<p>This is a sample app created using tuna-app code structure as the basic framework</p>

<h2> Prerequisites for Installing Hyperledger Frameworks</h2>
<ul>
  <li>cURL</li> 
  <li>Node.js</li>
  <li>npm package manager</li>
  <li>Go Language</li>
  <li>Docker</li>
  <li>Docker Compose</li> 
</ul>
<h2> Installation for MacOS</h2>
<ul>
  <li><b>curl -V</b></li> 
  <li>Check <b>https://nodejs.org/en/download/</b> for installing Node and NPM</li>
  <li>Installing GO
    <ul>
      <li><b>cd ~</b></li>
      <li><b>sudo curl -O https://storage.googleapis.com/golang/go1.9.2.darwin-amd64.pkg </b>(Check version at https://golang.org/dl/)</li>
      <li><b>open go1.9.2.darwin-amd64.pkg</b></li>
      <li><b>echo $GOPATH</b></li>
      <li>If nothing comes in the above step, go to .profile(or .bash_profile) and add 
        <ol>
          <li><b>export GOPATH=$HOME/go</b></li>
          <li><b>export PATH=$PATH:$GOPATH/bin</b></li>
        </ol>
        Then run source .profile to tun these new paths.
        </li>
      </ul>
  </li>
  <li>Visit <b>https://www.docker.com/docker-mac</b> to install docker for mac (Installing docker for mac or docker toolbox will also download docker-compose)</li>
</ul>

<h2>Installing hyperledger Fabric</h2>
<ul>
  <li><b>curl -sSL https://goo.gl/Q3YRTi | bash </b>(Check Check https://hyperledger-fabric.readthedocs.io/en/latest/samples.html#binaries for latest URL)</li>
  <li>docker images</li>
  <li>If the images are not in latest tag run <b>docker tag hyperledger/fabric-tools:x86_64-1.0.2 hyperledger/fabric-tools:latest</b> (Replace fabric-tools:x86_64-1.0.2 accordingly)</li>
  <li><b>export PATH=$PWD/bin:$PATH</b></li>
</ul>

<h2>Installing hyperledger Fabric Sample Code</h2>
<ul>
  <li><b>git clone https://github.com/hyperledger/fabric-samples.git</b></li>
</ul>

<h2>Getting started with the network</h2>
<ul>
  <li><b>cd fabric-samples/first-network</b></li>
  <li><b>./byfn.sh -m generate</b>
  <ul>
    <li>It generates all the certificates, keys and Genesis block required for network configuration</li>
  </ul>
  </li>
  <li><b>./byfn.sh -m up</b>
    <ul>
      <li>It starts the network</li>
    </ul>
  </li>
  <li><b>./byfn.sh -m down</b>
    <ul>
      <li>It shuts down the network</li>
    </ul>
  </li>
</ul>

<h2>Writing a sample application</h2>
<ul>
  <li>Download the education repository</li>
  <li><b>git clone https://github.com/hyperledger/education.git</b></li>
  <li><b>cd education/LFS171x/fabric-material/tuna-app</b></li>
  <li><b>docker rm -f $(docker ps -aq)</b> (To remove pre-existing containers)</li>
  <li><b>./startFabric.sh</b> (To start the network)</li>
  <li><b>npm install</b> (To install required packages from package.json file)</li>
  <li><b>node registerAdmin.js</b> (Register the admin)</li>
  <li><b>node registerUser.js</b> (Register the user)</li>
  <li><b>node server.js</b> (Start the application)</li>
  <li>open <b>localhost:8000</b> and your app would be running</li>
</ul>

<h2>Deploying the modified codechain</h2>
<ul>
  <li>Move your chaincode (i.e. invoice-chaincode.go) to /opt/gopath/src/github.com/invoice-app (Check the name in startFabric.sh and accordingly replace invoice-app)</li>
  <li>If the above directory is not present, then create the directoty and move</li>
  <li>Troubleshooting
    <ul>If <b>node registerUser.js</b>  gives error, run <b>rm -rf .hfc-key-store/ </b> and then run the command for both admin and user in order.</ul>
  </li>
</ul>


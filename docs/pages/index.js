import Head from 'next/head'
import Footer  from '../components/footer'
import Navbar from '../components/navbar'
import Section from '../components/section'

const Home = () => {
  return (
    <div className="bg-gray-200">
      <Head>
        <title>lfl - Links for lectures</title>
      </Head>
      <main>
        <Navbar/>
        <div className="container px-4 lg:px-40 mx-auto pt-16">
          <div>
            <Section
              id="#demonstation"
              title="Demonstation"
              subtitle="Preview of the main functionallity of lfl"
              description="lfl - Links for Lectures is a command line interface which opens the lectures' virtual meeting of your choice
              from terminal. It currently supports only Zoom and Webex links."
            >
              <div className="mt-6">
                <img src="/preview.png"/>
              </div>
            </Section>

             <Section
              id="#installation"
              title="Installation"
              subtitle="Preview of the installation of lfl"
              description="lfl is tested only in Linux environments. Please, use the instructions below to install it on your system."
            >

            <h3 className="font-bold text-xl mt-6">Debian & Debian-Based distributions</h3>
            <div className="mt-2">
              <img src="/debian-installation.png"/>
            </div>

            <h3 className="font-bold text-xl mt-6">Arch & Arch-Based distributions</h3>
            <div className="mt-2">
              <img src="/arch-installation.png"/>
            </div>

           <div className="mt-8">
              <a href="https://github.com/Stavrospanakakis/lfl" 
               className="px-4 py-2 rounded-md border-2 border-gray-900 bg-gray-900 text-white"
               target="_blank"
              >
                Source code
              </a>
            </div>
            </Section>

            <Section
              id="#usage"
              title="Usage"
              subtitle="Preview of the functionallity of lfl"
            >

              <h3 className="font-bold text-xl mt-6">Manual</h3>
              <div className="mt-2">
                <img className="lg:h-2/6 lg:w-4/6" src="/manual.png"/>
              </div>

              <h3 className="font-bold text-xl mt-6">First time running lfl</h3>
              <div className="mt-2">
                <img className="lg:h-2/6 lg:w-4/6" src="/first-time.png"/>
              </div>


              <h3 className="font-bold text-xl mt-6">Add lectures</h3>
              <div className="mt-2">
                <img className="lg:h-2/6 lg:w-4/6" src="/add.png"/>
              </div>


              <h3 className="font-bold text-xl mt-6">Remove lectures</h3>
              <div className="mt-2">
                <img className="lg:h-2/6 lg:w-4/6" src="/remove.png"/>
              </div>
            </Section>
          </div>
        </div>
        </main>
        <Footer/>
      </div>
      
  )
}

export default Home
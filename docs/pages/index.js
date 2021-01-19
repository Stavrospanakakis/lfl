import Head from 'next/head'
import Section from '../components/section'

export default function Home() {
  return (
    <div className="bg-gray-200">
      <Head>
        <title>Create Next App</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main>
        <header>
          <nav class="relative flex flex-wrap items-center justify-between px-2 py-3 navbar-expand-lg bg-black mb-3">
            <div class="container px-4 mx-auto flex flex-wrap items-center justify-between">
              <div class="w-full relative flex justify-between lg:w-auto  px-4 lg:static lg:block lg:justify-start">
                <a class="text-xl font-bold leading-relaxed inline-block mr-4 py-2 whitespace-no-wrap text-white" href="#pablo">
                  lfl - Links for Lectures
                </a>
                <button class="cursor-pointer text-xl leading-none px-3 py-1 border border-solid border-transparent rounded bg-transparent block lg:hidden outline-none focus:outline-none" type="button">
                  <span class="block relative w-6 h-px rounded-sm bg-white"></span>
                  <span class="block relative w-6 h-px rounded-sm bg-white mt-1"></span>
                  <span class="block relative w-6 h-px rounded-sm bg-white mt-1"></span>
                </button>
              </div>
              <div class="lg:flex flex-grow items-center" id="example-navbar-warning">
                <ul class="flex flex-col lg:flex-row list-none ml-auto">
                    <li class="nav-item">
                      <a class="px-3 py-2 flex items-center text-sm uppercase font-bold leading-snug text-white hover:opacity-75" href="#installation">
                        Installation
                      </a>
                    </li>
                    <li class="nav-item">
                      <a class="px-3 py-2 flex items-center text-sm uppercase font-bold leading-snug text-white hover:opacity-75" href="#usage">
                        Usage
                      </a>
                    </li>
                </ul>
              </div>
            </div>
          </nav>
        </header>
        <div className="container px-4 lg:px-40 mx-auto">
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
              <img className="lg:h-2/6 lg:w-4/6" src="/arch-installation.png"/>
            </div>

            <h3 className="font-bold text-xl mt-6">Arch & Arch-Based distributions</h3>
            <div className="mt-2">
              <img className="lg:h-2/6 lg:w-4/6" src="/arch-installation.png"/>
            </div>

            <h3 className="font-bold text-xl mt-6">Other distributions</h3>
            <div className="mt-2">
              <img className="lg:h-2/6 lg:w-4/6" src="/arch-installation.png"/>
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
        <footer>
              <div className="bg-black border border-gray-300 mt-10"></div>
              <div align="center" className="mt-2 p-5">© 2021 Stavros Panakakis</div>
        </footer>
      </div>
  )
}

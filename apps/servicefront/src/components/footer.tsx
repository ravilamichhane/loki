import { Link } from "@tanstack/react-router";

export function Footer() {
  return (
    <footer key="1" className="bg-gray-900 text-gray-400 py-12 md:py-16">
      <div className="container mx-auto px-4 md:px-6 lg:px-8">
        <div className="grid grid-cols-1 md:grid-cols-5 gap-8 md:gap-12">
          <div className="flex flex-col items-start justify-between h-full">
            <Link className="inline-flex items-center" href="#">
              <FlagIcon className="h-8 w-auto mr-2 text-white" />
              <span className="text-xl font-bold text-white">Ecommerce Store</span>
            </Link>
            <div className="mt-auto">
              <p className="text-gray-400 text-sm">© 2024 Ecommerce Store. All rights reserved.</p>
            </div>
          </div>
          <div>
            <h3 className="text-white font-bold mb-4 md:mb-6">Information</h3>
            <nav className="flex justify-start">
              <ul className="grid grid-cols-1 gap-4 md:gap-6">
                <li>
                  <Link className="hover:text-white transition-colors" href="#">
                    About
                  </Link>
                </li>
                <li>
                  <Link className="hover:text-white transition-colors" href="#">
                    Contact
                  </Link>
                </li>
                <li>
                  <Link className="hover:text-white transition-colors" href="#">
                    Shipping
                  </Link>
                </li>
                <li>
                  <Link className="hover:text-white transition-colors" href="#">
                    Returns
                  </Link>
                </li>
              </ul>
            </nav>
          </div>
          <div>
            <h3 className="text-white font-bold mb-4 md:mb-6">Support</h3>
            <nav className="flex justify-start">
              <ul className="grid grid-cols-1 gap-4 md:gap-6">
                <li>
                  <Link className="hover:text-white transition-colors" href="#">
                    Help Center
                  </Link>
                </li>
                <li>
                  <Link className="hover:text-white transition-colors" href="#">
                    Live Chat
                  </Link>
                </li>
                <li>
                  <Link className="hover:text-white transition-colors" href="#">
                    Email Support
                  </Link>
                </li>
                <li>
                  <Link className="hover:text-white transition-colors" href="#">
                    Community Forum
                  </Link>
                </li>
              </ul>
            </nav>
          </div>
          <div>
            <h3 className="text-white font-bold mb-4 md:mb-6">Developers</h3>
            <nav className="flex justify-start">
              <ul className="grid grid-cols-1 gap-4 md:gap-6">
                <li>
                  <Link className="hover:text-white transition-colors" href="#">
                    API Documentation
                  </Link>
                </li>
                <li>
                  <Link className="hover:text-white transition-colors" href="#">
                    SDK Downloads
                  </Link>
                </li>
                <li>
                  <Link className="hover:text-white transition-colors" href="#">
                    Developer Tools
                  </Link>
                </li>
                <li>
                  <Link className="hover:text-white transition-colors" href="#">
                    Developer Community
                  </Link>
                </li>
              </ul>
            </nav>
          </div>
          <div>
            <h3 className="text-white font-bold mb-4 md:mb-6">Products</h3>
            <nav className="flex justify-start">
              <ul className="grid grid-cols-1 gap-4 md:gap-6">
                <li>
                  <Link className="hover:text-white transition-colors" href="#">
                    Apparel
                  </Link>
                </li>
                <li>
                  <Link className="hover:text-white transition-colors" href="#">
                    Accessories
                  </Link>
                </li>
                <li>
                  <Link className="hover:text-white transition-colors" href="#">
                    Home Goods
                  </Link>
                </li>
                <li>
                  <Link className="hover:text-white transition-colors" href="#">
                    Electronics
                  </Link>
                </li>
              </ul>
            </nav>
          </div>
        </div>
        <div className="mt-12 md:mt-16 border-t border-gray-800 pt-8 md:pt-12">
          <div className="flex flex-col md:flex-row justify-between items-center">
            <div className="flex space-x-4">
              <Link className="hover:text-white transition-colors" href="#">
                FAQ
              </Link>
              <Link className="hover:text-white transition-colors" href="#">
                Privacy Policy
              </Link>
              <Link className="hover:text-white transition-colors" href="#">
                Terms
              </Link>
              <Link className="hover:text-white transition-colors" href="#">
                Careers
              </Link>
            </div>
            <div className="flex space-x-4 mt-4 md:mt-0">
              <Link className="hover:text-white transition-colors" href="#">
                <FacebookIcon className="h-5 w-5" />
              </Link>
              <Link className="hover:text-white transition-colors" href="#">
                <TwitterIcon className="h-5 w-5" />
              </Link>
              <Link className="hover:text-white transition-colors" href="#">
                <InstagramIcon className="h-5 w-5" />
              </Link>
              <Link className="hover:text-white transition-colors" href="#">
                <LinkedinIcon className="h-5 w-5" />
              </Link>
            </div>
          </div>
        </div>
      </div>
    </footer>
  );
}

function FacebookIcon(props: any) {
  return (
    <svg
      {...props}
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      <path d="M18 2h-3a5 5 0 0 0-5 5v3H7v4h3v8h4v-8h3l1-4h-4V7a1 1 0 0 1 1-1h3z" />
    </svg>
  );
}

function FlagIcon(props: any) {
  return (
    <svg
      {...props}
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      <path d="M4 15s1-1 4-1 5 2 8 2 4-1 4-1V3s-1 1-4 1-5-2-8-2-4 1-4 1z" />
      <line x1="4" x2="4" y1="22" y2="15" />
    </svg>
  );
}

function InstagramIcon(props: any) {
  return (
    <svg
      {...props}
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      <rect width="20" height="20" x="2" y="2" rx="5" ry="5" />
      <path d="M16 11.37A4 4 0 1 1 12.63 8 4 4 0 0 1 16 11.37z" />
      <line x1="17.5" x2="17.51" y1="6.5" y2="6.5" />
    </svg>
  );
}

function LinkedinIcon(props: any) {
  return (
    <svg
      {...props}
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      <path d="M16 8a6 6 0 0 1 6 6v7h-4v-7a2 2 0 0 0-2-2 2 2 0 0 0-2 2v7h-4v-7a6 6 0 0 1 6-6z" />
      <rect width="4" height="12" x="2" y="9" />
      <circle cx="4" cy="4" r="2" />
    </svg>
  );
}

function TwitterIcon(props: any) {
  return (
    <svg
      {...props}
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      <path d="M22 4s-.7 2.1-2 3.4c1.6 10-9.4 17.3-18 11.6 2.2.1 4.4-.6 6-2C3 15.5.5 9.6 3 5c2.2 2.6 5.6 4.1 9 4-.9-4.2 4-6.6 7-3.8 1.1 0 3-1.2 3-1.2z" />
    </svg>
  );
}

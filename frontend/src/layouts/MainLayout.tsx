import { Fragment } from 'react';
import { Outlet, Link } from 'react-router-dom';
import { Disclosure, Menu, Transition } from '@headlessui/react';
import { Bars3Icon, XMarkIcon, UserCircleIcon } from '@heroicons/react/24/outline';
import { useSelector } from 'react-redux';
import { RootState } from '../store';

const navigation = [
  { name: 'Home', href: '/' },
  { name: 'Cars', href: '/cars' },
  { name: 'Blog', href: '/blog' },
];

const MainLayout = () => {
  const { isAuthenticated, user } = useSelector((state: RootState) => state.auth);

  return (
    <div className="min-h-screen bg-gray-100">
      <Disclosure as="nav" className="bg-white shadow-sm">
        {({ open }) => (
          <>
            <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
              <div className="flex h-16 justify-between">
                <div className="flex">
                  <div className="flex flex-shrink-0 items-center">
                    <Link to="/" className="text-2xl font-bold text-primary-600">
                      CarRental
                    </Link>
                  </div>
                  <div className="hidden sm:ml-6 sm:flex sm:space-x-8">
                    {navigation.map((item) => (
                      <Link
                        key={item.name}
                        to={item.href}
                        className="inline-flex items-center border-b-2 border-transparent px-1 pt-1 text-sm font-medium text-gray-500 hover:border-gray-300 hover:text-gray-700"
                      >
                        {item.name}
                      </Link>
                    ))}
                  </div>
                </div>
                <div className="hidden sm:ml-6 sm:flex sm:items-center">
                  {isAuthenticated ? (
                    <Menu as="div" className="relative ml-3">
                      <Menu.Button className="flex rounded-full bg-white text-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:ring-offset-2">
                        <UserCircleIcon className="h-8 w-8 text-gray-400" />
                      </Menu.Button>
                      <Transition
                        as={Fragment}
                        enter="transition ease-out duration-200"
                        enterFrom="transform opacity-0 scale-95"
                        enterTo="transform opacity-100 scale-100"
                        leave="transition ease-in duration-75"
                        leaveFrom="transform opacity-100 scale-100"
                        leaveTo="transform opacity-0 scale-95"
                      >
                        <Menu.Items className="absolute right-0 z-10 mt-2 w-48 origin-top-right rounded-md bg-white py-1 shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none">
                          {user?.roles.includes('admin') && (
                            <Menu.Item>
                              {({ active }) => (
                                <Link
                                  to="/admin"
                                  className={`block px-4 py-2 text-sm text-gray-700 ${
                                    active ? 'bg-gray-100' : ''
                                  }`}
                                >
                                  Admin Dashboard
                                </Link>
                              )}
                            </Menu.Item>
                          )}
                          <Menu.Item>
                            {({ active }) => (
                              <Link
                                to="/profile"
                                className={`block px-4 py-2 text-sm text-gray-700 ${
                                  active ? 'bg-gray-100' : ''
                                }`}
                              >
                                Your Profile
                              </Link>
                            )}
                          </Menu.Item>
                          <Menu.Item>
                            {({ active }) => (
                              <button
                                className={`block w-full px-4 py-2 text-left text-sm text-gray-700 ${
                                  active ? 'bg-gray-100' : ''
                                }`}
                              >
                                Sign out
                              </button>
                            )}
                          </Menu.Item>
                        </Menu.Items>
                      </Transition>
                    </Menu>
                  ) : (
                    <div className="space-x-4">
                      <Link
                        to="/login"
                        className="rounded-md bg-white px-3 py-2 text-sm font-semibold text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 hover:bg-gray-50"
                      >
                        Sign in
                      </Link>
                      <Link
                        to="/register"
                        className="rounded-md bg-primary-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-primary-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primary-600"
                      >
                        Sign up
                      </Link>
                    </div>
                  )}
                </div>
                <div className="-mr-2 flex items-center sm:hidden">
                  <Disclosure.Button className="inline-flex items-center justify-center rounded-md p-2 text-gray-400 hover:bg-gray-100 hover:text-gray-500 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-primary-500">
                    <span className="sr-only">Open main menu</span>
                    {open ? (
                      <XMarkIcon className="block h-6 w-6" aria-hidden="true" />
                    ) : (
                      <Bars3Icon className="block h-6 w-6" aria-hidden="true" />
                    )}
                  </Disclosure.Button>
                </div>
              </div>
            </div>

            <Disclosure.Panel className="sm:hidden">
              <div className="space-y-1 pb-3 pt-2">
                {navigation.map((item) => (
                  <Link
                    key={item.name}
                    to={item.href}
                    className="block border-l-4 border-transparent py-2 pl-3 pr-4 text-base font-medium text-gray-500 hover:border-gray-300 hover:bg-gray-50 hover:text-gray-700"
                  >
                    {item.name}
                  </Link>
                ))}
              </div>
              {!isAuthenticated && (
                <div className="border-t border-gray-200 pb-3 pt-4">
                  <div className="space-y-1">
                    <Link
                      to="/login"
                      className="block px-4 py-2 text-base font-medium text-gray-500 hover:bg-gray-100 hover:text-gray-800"
                    >
                      Sign in
                    </Link>
                    <Link
                      to="/register"
                      className="block px-4 py-2 text-base font-medium text-gray-500 hover:bg-gray-100 hover:text-gray-800"
                    >
                      Sign up
                    </Link>
                  </div>
                </div>
              )}
            </Disclosure.Panel>
          </>
        )}
      </Disclosure>

      <main>
        <div className="mx-auto max-w-7xl py-6 sm:px-6 lg:px-8">
          <Outlet />
        </div>
      </main>

      <footer className="bg-white border-t">
        <div className="container mx-auto px-4 py-12">
          <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
            {/* About Us Column */}
            <div>
              <h5 className="text-xl font-semibold mb-6 text-gray-900">About us</h5>
              <div className="text-gray-600 space-y-2">
                <p>Car rental company</p>
                <p>1234 Main Street</p>
                <p>Los Angeles, CA 90001</p>
                <p>United States</p>
              </div>
              <div className="mt-6 space-y-3">
                <a
                  href="tel:609544872"
                  className="flex items-center gap-2 px-4 py-2 bg-gray-50 text-gray-700 rounded-2xl border border-gray-200 hover:bg-gray-100 transition-colors w-fit"
                >
                  <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M3 5a2 2 0 012-2h3.28a1 1 0 01.948.684l1.498 4.493a1 1 0 01-.502 1.21l-2.257 1.13a11.042 11.042 0 005.516 5.516l1.13-2.257a1 1 0 011.21-.502l4.493 1.498a1 1 0 01.684.949V19a2 2 0 01-2 2h-1C9.716 21 3 14.284 3 6V5z" />
                  </svg>
                  609544872
                </a>
                <a
                  href="mailto:car.rental@gmail.com"
                  className="flex items-center gap-2 px-4 py-2 bg-gray-50 text-gray-700 rounded-2xl border border-gray-200 hover:bg-gray-100 transition-colors w-fit"
                >
                  <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
                  </svg>
                  car.rental@gmail.com
                </a>
              </div>
            </div>

            {/* Our Partners Column */}
            <div>
              <h5 className="text-xl font-semibold mb-6 text-center text-gray-900">Our Partners</h5>
              <div className="grid grid-cols-2 gap-4">
                {[
                  'https://res.cloudinary.com/dpg94mnti/image/upload/v1685379456/obraz_2023-05-29_185733156_dyfrbc.png',
                  'https://res.cloudinary.com/dpg94mnti/image/upload/v1685379535/obraz_2023-05-29_185852827_qs7nxt.png',
                  'https://res.cloudinary.com/dpg94mnti/image/upload/v1685379643/obraz_2023-05-29_190040932_rjkoxy.png',
                  'https://res.cloudinary.com/dpg94mnti/image/upload/v1685379967/obraz_2023-05-29_190605002_shsxwk.png',
                  'https://res.cloudinary.com/dpg94mnti/image/upload/v1685380423/obraz_2023-05-29_191341300_hvjve2.png',
                  'https://res.cloudinary.com/dpg94mnti/image/upload/v1685379725/obraz_2023-05-29_190202721_xua3ya.png',
                ].map((src, index) => (
                  <div key={index} className={`flex ${index % 2 === 0 ? 'justify-end' : 'justify-start'}`}>
                    <a href="#" className="hover:opacity-80 transition-opacity">
                      <img src={src} alt={`Partner ${index + 1}`} className="w-24 h-auto" />
                    </a>
                  </div>
                ))}
              </div>
            </div>

            {/* Social Media Column */}
            <div>
              <h5 className="text-xl font-semibold mb-6 text-right text-gray-900">Social Media</h5>
              <div className="space-y-4">
                <a href="#" className="flex items-center justify-end gap-2 text-gray-700 hover:text-gray-900 transition-colors text-lg">
                  <span>facebook</span>
                  <svg className="h-6 w-6" fill="currentColor" viewBox="0 0 24 24">
                    <path d="M18.77,7.46H14.5v-1.9c0-.9.6-1.1,1-1.1h3V.5h-4.33C10.24.5,9.5,3.44,9.5,5.32v2.15h-3v4h3v12h5v-12h3.85l.42-4Z"/>
                  </svg>
                </a>
                <a href="#" className="flex items-center justify-end gap-2 text-gray-700 hover:text-gray-900 transition-colors text-lg">
                  <span>instagram</span>
                  <svg className="h-6 w-6" fill="currentColor" viewBox="0 0 24 24">
                    <path d="M12,2.2c3.2,0,3.6,0,4.9.1,3.3.1,4.8,1.7,4.9,4.9.1,1.3.1,1.6.1,4.8,0,3.2,0,3.6-.1,4.8-.1,3.2-1.7,4.8-4.9,4.9-1.3.1-1.6.1-4.9.1-3.2,0-3.6,0-4.8-.1-3.3-.1-4.8-1.7-4.9-4.9-.1-1.3-.1-1.6-.1-4.8,0-3.2,0-3.6.1-4.8C2.4,4,4,2.4,7.2,2.3,8.5,2.2,8.8,2.2,12,2.2ZM12,0C8.7,0,8.3,0,7.1.1,2.7.3.3,2.7.1,7.1,0,8.3,0,8.7,0,12s0,3.7.1,4.9c.2,4.4,2.6,6.8,7,7,1.2.1,1.6.1,4.9.1s3.7,0,4.9-.1c4.4-.2,6.8-2.6,7-7,.1-1.2.1-1.6.1-4.9s0-3.7-.1-4.9c-.2-4.4-2.6-6.8-7-7C15.7,0,15.3,0,12,0Zm0,5.8A6.2,6.2,0,1,0,18.2,12,6.2,6.2,0,0,0,12,5.8Zm0,10.2A4,4,0,1,1,16,12,4,4,0,0,1,12,16Zm7.8-10.4a1.4,1.4,0,1,1-1.4-1.4A1.4,1.4,0,0,1,19.8,5.6Z"/>
                  </svg>
                </a>
                <a href="#" className="flex items-center justify-end gap-2 text-gray-700 hover:text-gray-900 transition-colors text-lg">
                  <span>twitter</span>
                  <svg className="h-6 w-6" fill="currentColor" viewBox="0 0 24 24">
                    <path d="M23.954 4.569c-.885.389-1.83.654-2.825.775 1.014-.611 1.794-1.574 2.163-2.723-.951.555-2.005.959-3.127 1.184-.896-.959-2.173-1.559-3.591-1.559-2.717 0-4.92 2.203-4.92 4.917 0 .39.045.765.127 1.124C7.691 8.094 4.066 6.13 1.64 3.161c-.427.722-.666 1.561-.666 2.475 0 1.71.87 3.213 2.188 4.096-.807-.026-1.566-.248-2.228-.616v.061c0 2.385 1.693 4.374 3.946 4.827-.413.111-.849.171-1.296.171-.314 0-.615-.03-.916-.086.631 1.953 2.445 3.377 4.604 3.417-1.68 1.319-3.809 2.105-6.102 2.105-.39 0-.779-.023-1.17-.067 2.189 1.394 4.768 2.209 7.557 2.209 9.054 0 13.999-7.496 13.999-13.986 0-.209 0-.42-.015-.63.961-.689 1.8-1.56 2.46-2.548l-.047-.02z"/>
                  </svg>
                </a>
              </div>
            </div>
          </div>
        </div>

        {/* Bottom Section */}
        <div className="border-t">
          <div className="container mx-auto px-4 py-6">
            <div className="flex flex-col sm:flex-row justify-between items-center gap-4">
              <p className="text-gray-600 text-sm">
                &copy; {new Date().getFullYear()} CarRental. All rights reserved.
              </p>
              <div className="flex gap-6">
                <a href="#" className="text-gray-600 hover:text-gray-900 transition-colors">
                  <svg className="h-5 w-5" fill="currentColor" viewBox="0 0 24 24">
                    <path d="M23.954 4.569c-.885.389-1.83.654-2.825.775 1.014-.611 1.794-1.574 2.163-2.723-.951.555-2.005.959-3.127 1.184-.896-.959-2.173-1.559-3.591-1.559-2.717 0-4.92 2.203-4.92 4.917 0 .39.045.765.127 1.124C7.691 8.094 4.066 6.13 1.64 3.161c-.427.722-.666 1.561-.666 2.475 0 1.71.87 3.213 2.188 4.096-.807-.026-1.566-.248-2.228-.616v.061c0 2.385 1.693 4.374 3.946 4.827-.413.111-.849.171-1.296.171-.314 0-.615-.03-.916-.086.631 1.953 2.445 3.377 4.604 3.417-1.68 1.319-3.809 2.105-6.102 2.105-.39 0-.779-.023-1.17-.067 2.189 1.394 4.768 2.209 7.557 2.209 9.054 0 13.999-7.496 13.999-13.986 0-.209 0-.42-.015-.63.961-.689 1.8-1.56 2.46-2.548l-.047-.02z"/>
                  </svg>
                </a>
                <a href="#" className="text-gray-600 hover:text-gray-900 transition-colors">
                  <svg className="h-5 w-5" fill="currentColor" viewBox="0 0 24 24">
                    <path d="M12,2.2c3.2,0,3.6,0,4.9.1,3.3.1,4.8,1.7,4.9,4.9.1,1.3.1,1.6.1,4.8,0,3.2,0,3.6-.1,4.8-.1,3.2-1.7,4.8-4.9,4.9-1.3.1-1.6.1-4.9.1-3.2,0-3.6,0-4.8-.1-3.3-.1-4.8-1.7-4.9-4.9-.1-1.3-.1-1.6-.1-4.8,0-3.2,0-3.6.1-4.8C2.4,4,4,2.4,7.2,2.3,8.5,2.2,8.8,2.2,12,2.2ZM12,0C8.7,0,8.3,0,7.1.1,2.7.3.3,2.7.1,7.1,0,8.3,0,8.7,0,12s0,3.7.1,4.9c.2,4.4,2.6,6.8,7,7,1.2.1,1.6.1,4.9.1s3.7,0,4.9-.1c4.4-.2,6.8-2.6,7-7,.1-1.2.1-1.6.1-4.9s0-3.7-.1-4.9c-.2-4.4-2.6-6.8-7-7C15.7,0,15.3,0,12,0Zm0,5.8A6.2,6.2,0,1,0,18.2,12,6.2,6.2,0,0,0,12,5.8Zm0,10.2A4,4,0,1,1,16,12,4,4,0,0,1,12,16Zm7.8-10.4a1.4,1.4,0,1,1-1.4-1.4A1.4,1.4,0,0,1,19.8,5.6Z"/>
                  </svg>
                </a>
                <a href="#" className="text-gray-600 hover:text-gray-900 transition-colors">
                  <svg className="h-5 w-5" fill="currentColor" viewBox="0 0 24 24">
                    <path d="M18.77,7.46H14.5v-1.9c0-.9.6-1.1,1-1.1h3V.5h-4.33C10.24.5,9.5,3.44,9.5,5.32v2.15h-3v4h3v12h5v-12h3.85l.42-4Z"/>
                  </svg>
                </a>
              </div>
            </div>
          </div>
        </div>
      </footer>
    </div>
  );
};

export default MainLayout; 
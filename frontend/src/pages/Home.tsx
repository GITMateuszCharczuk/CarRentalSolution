import { Link } from 'react-router-dom';

const Home = () => {
  return (
    <div className="space-y-16">
      {/* Hero Section */}
      <div className="relative">
        <div className="absolute inset-0">
          <div className="bg-gradient-to-r from-primary-600 to-primary-800 opacity-90 absolute inset-0" />
        </div>
        <div className="relative">
          <div className="max-w-7xl mx-auto py-24 px-4 sm:py-32 sm:px-6 lg:px-8">
            <h1 className="text-4xl font-extrabold tracking-tight text-white sm:text-5xl lg:text-6xl">
              Rent Your Dream Car Today
            </h1>
            <p className="mt-6 max-w-3xl text-xl text-white">
              Experience luxury and comfort with our premium car rental service. Choose from our wide range of vehicles for any occasion.
            </p>
            <div className="mt-10">
              <Link
                to="/cars"
                className="inline-flex items-center px-6 py-3 border border-transparent text-base font-medium rounded-md shadow-sm text-primary-700 bg-white hover:bg-primary-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary-500"
              >
                Browse Cars
              </Link>
            </div>
          </div>
        </div>
      </div>

      {/* Featured Cars Section */}
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="text-center">
          <h2 className="text-3xl font-extrabold tracking-tight text-gray-900 sm:text-4xl">
            Featured Cars
          </h2>
          <p className="mt-4 max-w-2xl mx-auto text-xl text-gray-500">
            Choose from our selection of premium vehicles
          </p>
        </div>

        {/* Featured Cars Grid - To be populated with actual data */}
        <div className="mt-12 grid gap-8 grid-cols-1 sm:grid-cols-2 lg:grid-cols-3">
          {[1, 2, 3].map((car) => (
            <div
              key={car}
              className="relative bg-white rounded-lg shadow-lg overflow-hidden"
            >
              <div className="h-48 bg-gray-200" />
              <div className="p-6">
                <h3 className="text-lg font-semibold text-gray-900">Car Name</h3>
                <p className="mt-2 text-gray-500">Starting from $XX/day</p>
                <div className="mt-4">
                  <Link
                    to={`/cars/${car}`}
                    className="text-primary-600 hover:text-primary-500"
                  >
                    View Details →
                  </Link>
                </div>
              </div>
            </div>
          ))}
        </div>
      </div>

      {/* Why Choose Us Section */}
      <div className="bg-gray-50 py-16">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center">
            <h2 className="text-3xl font-extrabold text-gray-900">
              Why Choose Us
            </h2>
          </div>
          <div className="mt-10">
            <div className="grid grid-cols-1 gap-8 sm:grid-cols-2 lg:grid-cols-3">
              {[
                {
                  title: '24/7 Support',
                  description:
                    'Round-the-clock customer support for your convenience',
                },
                {
                  title: 'Flexible Rental',
                  description:
                    'Choose from hourly, daily, weekly, or monthly rentals',
                },
                {
                  title: 'Best Prices',
                  description:
                    'Competitive prices with no hidden fees or charges',
                },
              ].map((feature) => (
                <div
                  key={feature.title}
                  className="bg-white p-6 rounded-lg shadow-md"
                >
                  <h3 className="text-lg font-medium text-gray-900">
                    {feature.title}
                  </h3>
                  <p className="mt-2 text-gray-500">{feature.description}</p>
                </div>
              ))}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Home; 
import { Link } from 'react-router-dom';
import { useQuery } from '@tanstack/react-query';
import { carService, blogService, fileService } from '../services/api';
import Slider from 'react-slick';
import 'slick-carousel/slick/slick.css';
import 'slick-carousel/slick/slick-theme.css';

const styles = `
  .no-scrollbar::-webkit-scrollbar {
    display: none;
  }
  .no-scrollbar {
    -ms-overflow-style: none;
    scrollbar-width: none;
  }
`;

const Home = () => {
  const { data: carOffers } = useQuery({
    queryKey: ['carOffers'],
    queryFn: () => carService.getCarOffers({
      page_size: 10,
      current_page: 1
    }),
  });

  const { data: blogPosts } = useQuery({
    queryKey: ['latestBlogPosts'],
    queryFn: () => blogService.getBlogPosts({
      page_size: 3,
      current_page: 1
    }),
  });

  const sliderSettings = {
    dots: false,
    infinite: false,
    speed: 500,
    slidesToShow: 3,
    slidesToScroll: 1,
    responsive: [
      {
        breakpoint: 1024,
        settings: {
          slidesToShow: 2,
        },
      },
      {
        breakpoint: 640,
        settings: {
          slidesToShow: 1,
        },
      },
    ],
  };

  const latestPosts = blogPosts?.Items?.slice(0, 3) || [];

  return (
    <>
      <style>{styles}</style>
      {/* Hero Section */}
      <div 
        className="relative bg-cover bg-center bg-no-repeat h-[600px] -mx-[calc(50vw-50%)] w-[100vw] -mt-[23px]"
        style={{
          backgroundImage: "url('https://res.cloudinary.com/dpg94mnti/image/upload/v1685991560/obraz_2023-06-05_205053192_wrykwi.png')",
        }}
      >
        <div className="absolute inset-0 bg-black/50" />
        <div className="relative h-full flex items-center">
          <div className="container mx-auto px-8">
            <div className="max-w-3xl">
              <h1 className="text-6xl font-bold text-white mb-8 [text-shadow:_1px_1px_0_rgb(0_0_0_/_100%)]">
                Start your own journey
              </h1>
              <p className="text-xl text-white mb-8 max-w-2xl">
                Embark on your adventure with our hassle-free car rental service, empowering you to start your own journey wherever the road may take you.
              </p>
            </div>
          </div>
        </div>
      </div>

      {/* New Cars Section */}
      <div className="container mx-auto px-4 py-16">
        <h2 className="text-3xl font-bold mb-8 text-gray-900">New stuff</h2>
        {carOffers?.Items && carOffers.Items.length > 0 ? (
          <div className="mb-8 overflow-x-auto no-scrollbar">
            <Slider {...sliderSettings} className="mx-[-1rem]">
              {carOffers.Items.map((car) => (
                <div key={car.id} className="px-4">
                  <div className="bg-white rounded-3xl shadow-lg overflow-hidden">
                    <div className="p-4">
                      <img
                        src={fileService.getFileUrl(car.featured_image_url)}
                        alt={car.heading}
                        className="w-full h-48 object-cover rounded-2xl"
                      />
                      <div className="p-4">
                        <h3 className="text-xl font-bold text-gray-900">{car.heading}</h3>
                        <p className="text-lg text-gray-600">{car.horsepower} hp</p>
                      </div>
                      <div className="bg-gray-50 p-4 rounded-xl">
                        <div className="flex justify-between items-center">
                          <p className="font-bold text-gray-900">
                            from ${car.one_normal_day_price} USD
                          </p>
                          <Link
                            to={`/cars/${car.id}`}
                            className="bg-primary-600 text-white p-2 rounded-xl hover:bg-primary-700 transition-colors"
                          >
                            <svg xmlns="http://www.w3.org/2000/svg" className="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M14 5l7 7m0 0l-7 7m7-7H3" />
                            </svg>
                          </Link>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              ))}
            </Slider>
          </div>
        ) : (
          <p>No cars available at the moment.</p>
        )}

        {/* Latest News Section */}
        <h2 className="text-3xl font-bold mb-8 mt-16 text-gray-900">Latest news</h2>
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
          {latestPosts.map((post) => (
            <div key={post.id} className="bg-white rounded-3xl shadow-lg overflow-hidden h-full flex flex-col">
              <div className="aspect-[16/9] overflow-hidden">
                <img
                  src={fileService.getFileUrl(post.featuredImageUrl)}
                  alt={post.heading}
                  className="w-full h-full object-cover"
                />
              </div>
              <div className="p-6 flex flex-col flex-grow">
                <h3 className="text-xl font-bold mb-4 text-gray-900">{post.heading}</h3>
                <p className="text-gray-600 mb-4 line-clamp-3">{post.shortDescription}</p>
              </div>
              <div className="p-4 flex justify-between items-center border-t">
                <span className="text-gray-600">
                  {new Date(post.publishedDate).toLocaleDateString()}
                </span>
                <Link
                  to={`/blog/${post.id}`}
                  className="text-primary-600 hover:text-primary-700 flex items-center gap-1"
                >
                  Read more
                  <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M14 5l7 7m0 0l-7 7m7-7H3" />
                  </svg>
                </Link>
              </div>
            </div>
          ))}
        </div>

        {/* About Us Section */}
        <h2 className="text-3xl font-bold mb-8 mt-16 text-gray-900">About us</h2>
        <div className="space-y-16">
          {/* First Person */}
          <div className="flex flex-col md:flex-row gap-8 items-center">
            <div className="md:w-1/3">
              <img
                src="https://res.cloudinary.com/dpg94mnti/image/upload/v1685568197/800px-Andrzej_Person_Kancelaria_Senatu_u84d1u.jpg"
                alt="CEO and Founder"
                className="w-full rounded-3xl"
              />
            </div>
            <div className="md:w-2/3">
              <h3 className="text-2xl font-bold mb-4">Seweryn Chrzęszczodyrski</h3>
              <p className="text-gray-600 mb-4">CEO and founder of Car Rental Company.</p>
              <p className="text-gray-600 mb-4">
                The hardworking entrepreneur tirelessly revolutionized the industry through unwavering dedication, innovation, and a relentless pursuit of excellence.
              </p>
              <p className="text-gray-600">
                In the summer season, it may happen that he will be the one renting you a car.
              </p>
            </div>
          </div>

          {/* Second Person */}
          <div className="flex flex-col md:flex-row-reverse gap-8 items-center">
            <div className="md:w-1/3">
              <img
                src="https://res.cloudinary.com/dpg94mnti/image/upload/v1685568197/photo-1500648767791-00dcc994a43e_c5d0u3.jpg"
                alt="Co-founder"
                className="w-full rounded-3xl"
              />
            </div>
            <div className="md:w-2/3">
              <h3 className="text-2xl font-bold mb-4">Prawdźimir Zasadowski</h3>
              <p className="text-gray-600 mb-4">Co-Founder of Car Rental Company.</p>
              <p className="text-gray-600 mb-4">
                The pillar of strength of the car renting company drives growth through strategic collaboration, innovative initiatives, and valuable partnerships, ensuring market expansion and a strong position in the competitive car rental industry.
              </p>
              <p className="text-gray-600">
                Brings innovation by introducing innovative solutions and ideas.
              </p>
            </div>
          </div>

          {/* Third Person */}
          <div className="flex flex-col md:flex-row gap-8 items-center">
            <div className="md:w-1/3">
              <img
                src="https://res.cloudinary.com/dpg94mnti/image/upload/v1685568197/istockphoto-123935971-170667a_nbwpfn.webp"
                alt="Main Technician"
                className="w-full rounded-3xl"
              />
            </div>
            <div className="md:w-2/3">
              <h3 className="text-2xl font-bold mb-4">Waldemar Stodoła</h3>
              <p className="text-gray-600 mb-4">Main technician and mechanic.</p>
              <p className="text-gray-600 mb-4">
                Qualified handyman that diligently maintains and ensures the optimal condition of vehicles, demonstrating expertise, attention to detail, and a strong commitment to providing safe and reliable transportation options.
              </p>
              <p className="text-gray-600">
                He probably changed oil in car that you will rent, several times.
              </p>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default Home; 
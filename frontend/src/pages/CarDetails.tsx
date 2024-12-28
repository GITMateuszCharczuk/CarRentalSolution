import { useState } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { useQuery } from '@tanstack/react-query';
import { carService, fileService } from '../services/api';
import { useSelector } from 'react-redux';
import { selectIsAuthenticated } from '../store/slices/authSlice';
import { CarOfferTag, CarOfferImage, ListResponse } from '../types/api';
import Slider from 'react-slick';
import 'slick-carousel/slick/slick.css';
import 'slick-carousel/slick/slick-theme.css';

// Add styles to hide scrollbar
const styles = `
  .no-scrollbar::-webkit-scrollbar {
    display: none;
  }
  .no-scrollbar {
    -ms-overflow-style: none;
    scrollbar-width: none;
  }
`;

interface TagsResponse {
  Items: CarOfferTag[];
}

const CarDetails = () => {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();
  const isAuthenticated = useSelector(selectIsAuthenticated);
  const [selectedImage, setSelectedImage] = useState<string | null>(null);
  const [reservation, setReservation] = useState({
    startDate: '',
    endDate: '',
  });

  const { data: car, isLoading, error } = useQuery({
    queryKey: ['carOffer', id],
    queryFn: () => carService.getCarOfferById(id!),
    enabled: !!id,
  });

  const { data: tagsData } = useQuery<TagsResponse>({
    queryKey: ['carOfferTags', id],
    queryFn: () => carService.getCarOfferTags(id!),
    enabled: !!id,
  });

  const { data: imagesData } = useQuery<ListResponse<CarOfferImage>>({
    queryKey: ['carOfferImages', id],
    queryFn: () => carService.getCarOfferImages(id!),
    enabled: !!id,
  });

  const handleReservation = (e: React.FormEvent) => {
    e.preventDefault();
    if (!isAuthenticated) {
      navigate('/login', { state: { from: `/cars/${id}` } });
      return;
    }
    navigate(`/car-order/${id}?startDate=${reservation.startDate}&endDate=${reservation.endDate}`);
  };

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
          slidesToShow: 3,
        },
      },
      {
        breakpoint: 768,
        settings: {
          slidesToShow: 2,
        },
      },
      {
        breakpoint: 480,
        settings: {
          slidesToShow: 1,
        },
      },
    ],
  };

  if (isLoading) {
    return (
      <div className="flex h-96 items-center justify-center">
        <div className="h-32 w-32 animate-spin rounded-full border-b-2 border-t-2 border-primary-500"></div>
      </div>
    );
  }

  if (error || !car) {
    return (
      <div className="flex h-96 items-center justify-center">
        <p className="text-red-500">Error loading car details. Please try again later.</p>
      </div>
    );
  }

  const mainImageUrl = selectedImage || (car.car_offer.featured_image_url ? fileService.getFileUrl(car.car_offer.featured_image_url) : '/placeholder-car.jpg');

  return (
    <>
      <style>{styles}</style>
      <div className="container mx-auto mt-5">
        <div className="flex flex-col lg:flex-row justify-center gap-8">
          {/* Left Column - Car Details */}
          <div className="lg:w-2/3">
            {/* Back Button and Title */}
            <div className="flex items-center gap-3 mb-2">
              <a href="/cars" className="text-gray-800 hover:text-gray-600">
                <svg xmlns="http://www.w3.org/2000/svg" className="h-10 w-10" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M10 19l-7-7m0 0l7-7m-7 7h18" />
                </svg>
              </a>
              <h1 className="text-3xl font-bold text-gray-900">{car.car_offer.heading}</h1>
            </div>

            {/* Tags Section */}
            {tagsData?.Items && tagsData.Items.length > 0 && (
              <div className="mb-4">
                <div className="flex flex-wrap gap-2">
                  {tagsData.Items.map((tag) => (
                    <span
                      key={tag.id}
                      className="inline-flex items-center rounded-full bg-primary-100 px-4 py-1.5 text-sm font-semibold text-primary-800 ring-1 ring-inset ring-primary-200 hover:bg-primary-200 transition-colors"
                    >
                      {tag.name}
                    </span>
                  ))}
                </div>
              </div>
            )}

            {/* Main Image */}
            <div className="mb-4">
              <img
                src={mainImageUrl}
                alt={car.car_offer.heading}
                className="w-full h-[500px] object-cover rounded-2xl"
              />
            </div>

            {/* Image Gallery */}
            {imagesData?.Items && imagesData.Items.length > 0 && (
              <div className="mb-8 overflow-x-auto no-scrollbar">
                <Slider {...sliderSettings} className="mx-[-0.5rem]">
                  {imagesData.Items.map((image) => (
                    <div key={image.id} className="px-2">
                      <div className="rounded-lg overflow-hidden">
                        <img
                          src={fileService.getFileUrl(image.url)}
                          alt={car.car_offer.heading}
                          className="h-24 w-full cursor-pointer object-cover hover:opacity-80 transition-opacity rounded-lg"
                          onClick={() => setSelectedImage(fileService.getFileUrl(image.url))}
                        />
                      </div>
                    </div>
                  ))}
                </Slider>
              </div>
            )}

            {/* Car Specifications */}
            <h2 className="text-2xl font-bold mb-4 text-gray-900">Description and Technical Data</h2>
            <div className="grid grid-cols-2 md:grid-cols-5 gap-4 mb-8">
              <div className="bg-gray-50 p-4 rounded-xl">
                <div className="flex justify-center mb-2">
                  <svg xmlns="http://www.w3.org/2000/svg" className="h-8 w-8 text-[#1f2937]" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
                  </svg>
                </div>
                <div className="text-center text-sm text-[#1f2937]">Year of Production</div>
                <div className="text-center font-bold text-[#1f2937]">{car.car_offer.year_of_production}</div>
              </div>

              <div className="bg-gray-50 p-4 rounded-xl">
                <div className="flex justify-center mb-2">
                  <svg xmlns="http://www.w3.org/2000/svg" className="h-8 w-8 text-[#1f2937]" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M13 10V3L4 14h7v7l9-11h-7z" />
                  </svg>
                </div>
                <div className="text-center text-sm text-[#1f2937]">Engine</div>
                <div className="text-center font-bold text-[#1f2937]">{car.car_offer.engine_details}</div>
              </div>

              <div className="bg-gray-50 p-4 rounded-xl">
                <div className="flex justify-center mb-2">
                  <svg xmlns="http://www.w3.org/2000/svg" className="h-8 w-8 text-[#1f2937]" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M13 10V3L4 14h7v7l9-11h-7z" />
                  </svg>
                </div>
                <div className="text-center text-sm text-[#1f2937]">Horsepower</div>
                <div className="text-center font-bold text-[#1f2937]">{car.car_offer.horsepower}</div>
              </div>

              <div className="bg-gray-50 p-4 rounded-xl">
                <div className="flex justify-center mb-2">
                  <svg xmlns="http://www.w3.org/2000/svg" className="h-8 w-8 text-[#1f2937]" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                  </svg>
                </div>
                <div className="text-center text-sm text-[#1f2937]">Gearbox</div>
                <div className="text-center font-bold text-[#1f2937]">{car.car_offer.gearbox_details}</div>
              </div>

              <div className="bg-gray-50 p-4 rounded-xl">
                <div className="flex justify-center mb-2">
                  <svg xmlns="http://www.w3.org/2000/svg" className="h-8 w-8 text-[#1f2937]" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 12a4 4 0 100-8 4 4 0 000 8z" />
                  </svg>
                </div>
                <div className="text-center text-sm text-[#1f2937]">Drive</div>
                <div className="text-center font-bold text-[#1f2937]">{car.car_offer.drive_details}</div>
              </div>
            </div>

            {/* Description */}
            <div className="border-t border-b py-4 mb-8">
              <p className="text-gray-700">{car.car_offer.short_description}</p>
            </div>

            {/* Pricing Table */}
            <div className="mb-8">
              <table className="min-w-full divide-y divide-gray-200">
                <thead className="bg-gray-800">
                  <tr>
                    <th scope="col" className="px-6 py-4 text-left text-sm font-medium text-white uppercase tracking-wider">Time</th>
                    <th scope="col" className="px-6 py-4 text-left text-sm font-medium text-white uppercase tracking-wider">Price</th>
                  </tr>
                </thead>
                <tbody className="bg-white divide-y divide-gray-200">
                  <tr>
                    <td className="px-6 py-4 whitespace-nowrap text-base text-[#1f2937]">Day (Mon. - Thu.)</td>
                    <td className="px-6 py-4 whitespace-nowrap text-base font-medium text-[#1f2937]">${car.car_offer.one_normal_day_price}</td>
                  </tr>
                  <tr>
                    <td className="px-6 py-4 whitespace-nowrap text-base text-[#1f2937]">Day (Fri. - Sun.)</td>
                    <td className="px-6 py-4 whitespace-nowrap text-base font-medium text-[#1f2937]">${car.car_offer.one_weekend_day_price}</td>
                  </tr>
                  <tr>
                    <td className="px-6 py-4 whitespace-nowrap text-base text-[#1f2937]">Week</td>
                    <td className="px-6 py-4 whitespace-nowrap text-base font-medium text-[#1f2937]">${car.car_offer.one_week_price}</td>
                  </tr>
                  <tr>
                    <td className="px-6 py-4 whitespace-nowrap text-base text-[#1f2937]">Month</td>
                    <td className="px-6 py-4 whitespace-nowrap text-base font-medium text-[#1f2937]">${car.car_offer.one_month_price}</td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>

          {/* Right Column - Reservation Form */}
          <div className="lg:w-1/3">
            <div className="bg-white p-6 rounded-2xl border border-gray-800 shadow-lg">
              <form onSubmit={handleReservation} className="space-y-6">
                <div>
                  <h3 className="text-2xl font-bold mb-4 text-gray-900">Book Now</h3>
                  <h4 className="text-xl font-bold mb-4 text-gray-900">Availability</h4>
                </div>

                <div className="grid grid-cols-2 gap-4">
                  <div>
                    <label htmlFor="startDate" className="block text-sm font-medium text-gray-700">Start Date</label>
                    <input
                      type="datetime-local"
                      id="startDate"
                      required
                      className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 sm:text-sm text-gray-900"
                      value={reservation.startDate}
                      onChange={(e) => setReservation({ ...reservation, startDate: e.target.value })}
                    />
                  </div>
                  <div>
                    <label htmlFor="endDate" className="block text-sm font-medium text-gray-700">End Date</label>
                    <input
                      type="datetime-local"
                      id="endDate"
                      required
                      className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 sm:text-sm text-gray-900"
                      value={reservation.endDate}
                      onChange={(e) => setReservation({ ...reservation, endDate: e.target.value })}
                    />
                  </div>
                </div>

                <div>
                  <h4 className="text-xl font-bold mb-2 text-gray-900">General Terms and Conditions</h4>
                  <ul className="space-y-2 text-sm text-gray-600">
                    <li className="flex items-center gap-2">
                      <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5 text-primary-500" viewBox="0 0 20 20" fill="currentColor">
                        <path fillRule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clipRule="evenodd" />
                      </svg>
                      Security deposit - $1500
                    </li>
                    <li className="flex items-center gap-2">
                      <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5 text-primary-500" viewBox="0 0 20 20" fill="currentColor">
                        <path fillRule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clipRule="evenodd" />
                      </svg>
                      Daily mileage limit: 300 km
                    </li>
                    <li className="flex items-center gap-2">
                      <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5 text-primary-500" viewBox="0 0 20 20" fill="currentColor">
                        <path fillRule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clipRule="evenodd" />
                      </svg>
                      Fee for exceeding mileage: $2/km
                    </li>
                  </ul>
                </div>

                <button
                  type="submit"
                  className="w-full rounded-full bg-primary-600 px-4 py-2 text-base font-medium text-white shadow-sm hover:bg-primary-700 focus:outline-none focus:ring-2 focus:ring-primary-500 focus:ring-offset-2"
                >
                  Book Now
                </button>
              </form>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default CarDetails; 
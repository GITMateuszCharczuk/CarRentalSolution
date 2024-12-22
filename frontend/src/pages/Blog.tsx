import { useState } from 'react';
import { useQuery } from '@tanstack/react-query';
import { Link } from 'react-router-dom';
import { blogService } from '../services/api';

const Blog = () => {
  const [currentPage, setCurrentPage] = useState(1);

  const { data, isLoading, error } = useQuery({
    queryKey: ['blogPosts', currentPage],
    queryFn: () =>
      blogService.getBlogPosts({
        current_page: currentPage,
        page_size: 6,
      }),
  });

  if (isLoading) {
    return (
      <div className="flex h-96 items-center justify-center">
        <div className="h-32 w-32 animate-spin rounded-full border-b-2 border-t-2 border-primary-500"></div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="flex h-96 items-center justify-center">
        <p className="text-red-500">Error loading blog posts. Please try again later.</p>
      </div>
    );
  }

  return (
    <div className="space-y-8">
      {/* Header */}
      <div className="border-b border-gray-200 pb-5">
        <h1 className="text-3xl font-bold leading-tight text-gray-900">Blog</h1>
        <p className="mt-2 max-w-4xl text-sm text-gray-500">
          Stay updated with our latest news, tips, and insights about car rental and automotive
          industry.
        </p>
      </div>

      {/* Blog Posts Grid */}
      <div className="grid grid-cols-1 gap-8 sm:grid-cols-2 lg:grid-cols-3">
        {data?.items.map((post) => (
          <article
            key={post.id}
            className="flex flex-col overflow-hidden rounded-lg shadow-lg transition duration-300 hover:shadow-xl"
          >
            <div className="flex-shrink-0">
              <img
                className="h-48 w-full object-cover"
                src={post.featuredImageUrl}
                alt={post.heading}
              />
            </div>
            <div className="flex flex-1 flex-col justify-between bg-white p-6">
              <div className="flex-1">
                <p className="text-sm font-medium text-primary-600">
                  {post.tags.map((tag) => (
                    <span key={tag} className="mr-2">
                      #{tag}
                    </span>
                  ))}
                </p>
                <Link to={`/blog/${post.urlHandle}`} className="mt-2 block">
                  <p className="text-xl font-semibold text-gray-900">{post.heading}</p>
                  <p className="mt-3 text-base text-gray-500">{post.shortDescription}</p>
                </Link>
              </div>
              <div className="mt-6 flex items-center">
                <div className="flex-shrink-0">
                  <span className="sr-only">{post.author}</span>
                  <div className="h-10 w-10 rounded-full bg-gray-200"></div>
                </div>
                <div className="ml-3">
                  <p className="text-sm font-medium text-gray-900">{post.author}</p>
                  <div className="flex space-x-1 text-sm text-gray-500">
                    <time dateTime={post.publishedDate}>
                      {new Date(post.publishedDate).toLocaleDateString()}
                    </time>
                  </div>
                </div>
              </div>
            </div>
          </article>
        ))}
      </div>

      {/* Pagination */}
      {data && data.totalPages > 1 && (
        <div className="flex items-center justify-center space-x-2">
          <button
            onClick={() => setCurrentPage((prev) => Math.max(prev - 1, 1))}
            disabled={currentPage === 1}
            className="rounded-md bg-white px-3 py-2 text-sm font-semibold text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 hover:bg-gray-50 disabled:cursor-not-allowed disabled:opacity-50"
          >
            Previous
          </button>
          <span className="text-sm text-gray-700">
            Page {currentPage} of {data.totalPages}
          </span>
          <button
            onClick={() => setCurrentPage((prev) => Math.min(prev + 1, data.totalPages))}
            disabled={currentPage === data.totalPages}
            className="rounded-md bg-white px-3 py-2 text-sm font-semibold text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 hover:bg-gray-50 disabled:cursor-not-allowed disabled:opacity-50"
          >
            Next
          </button>
        </div>
      )}
    </div>
  );
};

export default Blog; 
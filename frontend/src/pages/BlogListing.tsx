import { useState, useMemo } from 'react';
import { Link } from 'react-router-dom';
import { useQuery } from '@tanstack/react-query';
import { blogService, fileService } from '../services/api';
import { Pagination } from '../components/Pagination';
import { SortSelect, type SortField } from '../components/SortSelect';
import { BlogPost, BlogTag } from '../types/api';
import { formatDateForApi, formatDateForInput } from '../utils/dateUtils';

const SORT_FIELDS: SortField[] = [
  { field: 'heading', label: 'Title' },
  { field: 'published_date', label: 'Published Date' },
  { field: 'likes_count', label: 'Likes' },
  { field: 'comments_count', label: 'Comments' },
];

const BlogListing = () => {
  const [currentPage, setCurrentPage] = useState(1);
  const [selectedTags, setSelectedTags] = useState<string[]>([]);
  const [sortFields, setSortFields] = useState<string[]>([]);
  const [filters, setFilters] = useState({
    dateTimeFrom: '',
    dateTimeTo: '',
  });

  // Validate and update dates when setting filters
  const handleDateChange = (field: 'dateTimeFrom' | 'dateTimeTo', value: string) => {
    // If setting end date and it's before start date, adjust it
    if (field === 'dateTimeTo' && filters.dateTimeFrom && value < filters.dateTimeFrom) {
      value = filters.dateTimeFrom;
    }

    // If setting start date and it's after end date, adjust end date
    if (field === 'dateTimeFrom' && filters.dateTimeTo && value > filters.dateTimeTo) {
      setFilters(prev => ({
        ...prev,
        [field]: value,
        dateTimeTo: value
      }));
      return;
    }

    setFilters(prev => ({
      ...prev,
      [field]: value
    }));
  };

  const { data: blogPosts, isLoading } = useQuery({
    queryKey: ['blogPosts', currentPage, selectedTags, sortFields, filters],
    queryFn: () => blogService.getBlogPosts({
      page_size: 9,
      current_page: currentPage,
      tags: selectedTags.length > 0 ? selectedTags : undefined,
      'date-time-from': filters.dateTimeFrom ? formatDateForApi(filters.dateTimeFrom) : undefined,
      'date-time-to': filters.dateTimeTo ? formatDateForApi(filters.dateTimeTo) : undefined,
      sort_fields: sortFields.length > 0 ? sortFields : undefined,
    }),
  });

  const { data: tags } = useQuery({
    queryKey: ['blogTags'],
    queryFn: () => blogService.getBlogTags(),
  });

  // Fetch likes and comments for all posts in a single query
  const postIds = blogPosts?.Items?.map(post => post.id) ?? [];
  
  const { data: likesData } = useQuery({
    queryKey: ['blogPostsLikes', postIds],
    queryFn: async () => {
      if (!postIds.length) return {};
      const likesPromises = postIds.map(id => blogService.getBlogPostLikes(id));
      const results = await Promise.all(likesPromises);
      return Object.fromEntries(postIds.map((id, index) => [id, results[index]]));
    },
    enabled: postIds.length > 0,
  });

  const { data: commentsData } = useQuery({
    queryKey: ['blogPostsComments', postIds],
    queryFn: async () => {
      if (!postIds.length) return {};
      const commentsPromises = postIds.map(id => blogService.getBlogPostCommentsCount(id));
      const results = await Promise.all(commentsPromises);
      return Object.fromEntries(postIds.map((id, index) => [id, results[index]]));
    },
    enabled: postIds.length > 0,
  });

  const handleTagClick = (tagName: string) => {
    setSelectedTags(prev => 
      prev.includes(tagName) 
        ? prev.filter(t => t !== tagName)
        : [...prev, tagName]
    );
    setCurrentPage(1);
  };

  if (isLoading) {
    return (
      <div className="flex h-96 items-center justify-center">
        <div className="h-32 w-32 animate-spin rounded-full border-b-2 border-t-2 border-primary-500"></div>
      </div>
    );
  }

  return (
    <div className="container mx-auto px-4">
      {/* Header Section */}
      <div className="mb-8">
        <h1 className="text-4xl font-bold text-gray-900 mb-2">Blog Posts</h1>
        <p className="text-lg text-gray-600">Stay updated with our latest news and stories</p>
      </div>

      {/* Filters Section */}
      <div className="bg-white rounded-3xl shadow-lg p-6 mb-8">
        <div className="flex justify-between items-center mb-4">
          <h2 className="text-xl font-semibold text-gray-900">Filter Posts</h2>
          <SortSelect
            availableFields={SORT_FIELDS}
            onChange={setSortFields}
            className="min-w-[200px] text-gray-900"
          />
        </div>
        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label htmlFor="dateFrom" className="block text-sm font-medium text-gray-700 mb-2">
              From Date
            </label>
            <input
              type="datetime-local"
              id="dateFrom"
              className="w-full rounded-xl border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 text-gray-900"
              value={filters.dateTimeFrom}
              onChange={(e) => handleDateChange('dateTimeFrom', e.target.value)}
            />
          </div>
          <div>
            <label htmlFor="dateTo" className="block text-sm font-medium text-gray-700 mb-2">
              To Date
            </label>
            <input
              type="datetime-local"
              id="dateTo"
              className="w-full rounded-xl border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 text-gray-900"
              value={filters.dateTimeTo}
              onChange={(e) => handleDateChange('dateTimeTo', e.target.value)}
            />
          </div>
        </div>
      </div>

      {/* Tags Section */}
      <div className="mb-8">
        <h2 className="text-xl font-semibold text-gray-900 mb-4">Categories</h2>
        <div className="flex flex-wrap gap-2">
          {tags?.Items && tags.Items.map((tag: BlogTag) => (
            <button
              key={tag.id}
              onClick={() => handleTagClick(tag.name)}
              className={`px-4 py-2 rounded-xl text-sm font-medium transition-all transform hover:scale-105
                ${selectedTags.includes(tag.name)
                  ? 'bg-gray-900 text-white shadow-lg'
                  : 'bg-white text-gray-900 shadow hover:shadow-md border border-gray-200'
                }`}
            >
              {tag.name}
            </button>
          ))}
        </div>
      </div>

      {/* Blog Posts Grid */}
      {blogPosts?.Items && blogPosts.Items.length > 0 ? (
        <div className="grid grid-cols-1 lg:grid-cols-3 md:grid-cols-2 gap-8">
          {blogPosts.Items.map((post: BlogPost) => (
            <div key={post.id} className="bg-white rounded-3xl shadow-lg overflow-hidden transform transition-all duration-300 hover:scale-[1.02] hover:shadow-xl flex flex-col">
              <div className="relative">
                <img
                  src={fileService.getFileUrl(post.featuredImageUrl)}
                  alt={post.heading}
                  className="w-full h-56 object-cover"
                />
                <div className="absolute inset-0 bg-gradient-to-t from-black/50 to-transparent"></div>
              </div>
              <div className="p-6 flex flex-col flex-1">
                <h3 className="text-2xl font-bold mb-4 text-gray-900 truncate" title={post.heading}>
                  {post.heading}
                </h3>
                <p className="text-gray-600 mb-4 line-clamp-3 flex-grow">{post.shortDescription}</p>
                
                <div className="flex justify-between items-center mt-auto">
                  <div className="flex items-center gap-4">
                    <div className="flex items-center text-gray-600">
                      <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" />
                      </svg>
                      <span>{likesData?.[post.id]?.TotalCount ?? 0}</span>
                    </div>
                    <div className="flex items-center text-gray-600">
                      <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M7 8h10M7 12h4m1 8l-4-4H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-3l-4 4z" />
                      </svg>
                      <span>{commentsData?.[post.id]?.Count ?? 0}</span>
                    </div>
                  </div>
                  <Link
                    to={`/blog/${post.id}`}
                    className="bg-gray-900 text-white py-2 px-4 rounded-xl text-sm font-medium hover:bg-gray-800 transition-colors"
                  >
                    Read More
                  </Link>
                </div>
              </div>
            </div>
          ))}
        </div>
      ) : (
        <div className="text-center py-12">
          <h3 className="text-xl font-medium text-gray-900 mb-2">No blog posts found</h3>
          <p className="text-gray-600">Try adjusting your search criteria</p>
        </div>
      )}

      {/* Pagination */}
      {blogPosts && blogPosts.TotalPages > 1 && (
        <div className="mt-12">
          <Pagination
            currentPage={currentPage}
            totalPages={blogPosts.TotalPages}
            onPageChange={setCurrentPage}
          />
        </div>
      )}
    </div>
  );
};

export default BlogListing; 
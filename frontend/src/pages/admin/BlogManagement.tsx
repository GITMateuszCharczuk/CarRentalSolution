import { useState } from 'react';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { blogService } from '../../services/api';
import type { BlogPost, BlogPostsQueryParams } from '../../types/api';
import { formatDateForApi } from '../../utils/dateUtils';
import { SortSelect, type SortField } from '../../components/SortSelect';
import { Pagination } from '../../components/Pagination';

const SORT_FIELDS: SortField[] = [
  { field: 'heading', label: 'Title' },
  { field: 'author', label: 'Author' },
  { field: 'publishedDate', label: 'Published Date' },
  { field: 'likes_count', label: 'Likes' },
  { field: 'comments_count', label: 'Comments' },
];

const BlogManagement = () => {
  const queryClient = useQueryClient();
  const [currentPage, setCurrentPage] = useState(1);
  const [pageSize, setPageSize] = useState(10);
  const [sortFields, setSortFields] = useState<string[]>([]);
  const [selectedTags, setSelectedTags] = useState<string[]>([]);
  const [dateFrom, setDateFrom] = useState<string>('');
  const [dateTo, setDateTo] = useState<string>('');
  const [visible, setVisible] = useState<boolean | undefined>(undefined);

  // Query for blog posts with pagination, sorting, and filtering
  const { data, isLoading } = useQuery({
    queryKey: ['adminBlogPosts', currentPage, pageSize, sortFields, selectedTags, dateFrom, dateTo, visible],
    queryFn: () => {
      const params: BlogPostsQueryParams = {
        current_page: currentPage,
        page_size: pageSize,
        sort_fields: sortFields.length > 0 ? sortFields : undefined,
        tags: selectedTags,
        'date-time-from': dateFrom ? formatDateForApi(dateFrom) : undefined,
        'date-time-to': dateTo ? formatDateForApi(dateTo) : undefined,
        visible: visible
      };
      return blogService.getBlogPosts(params);
    },
  });

  // Mutations for CRUD operations
  const deleteMutation = useMutation({
    mutationFn: (postId: string) => blogService.deleteBlogPost(postId),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['adminBlogPosts'] });
    },
  });

  const toggleVisibilityMutation = useMutation({
    mutationFn: async (post: BlogPost) => {
      const updateData = {
        ...post,
        visible: !post.visible,
      };
      return blogService.updateBlogPost(post.id, updateData);
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['adminBlogPosts'] });
    },
  });

  const handleDelete = async (postId: string) => {
    if (window.confirm('Are you sure you want to delete this blog post?')) {
      await deleteMutation.mutateAsync(postId);
    }
  };

  return (
    <div className="space-y-6">
      <div className="sm:flex sm:items-center sm:justify-between">
        <div>
          <h1 className="text-2xl font-bold text-gray-900">Blog Posts Management</h1>
          <p className="mt-2 text-sm text-gray-700">
            Manage your blog posts, including creation, editing, and deletion.
          </p>
        </div>
        <button
          type="button"
          onClick={() => {/* TODO: Implement create/edit modal */}}
          className="inline-flex items-center rounded-md bg-primary-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-primary-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primary-600"
        >
          Create New Post
        </button>
      </div>

      {/* Filters */}
      <div className="bg-white p-4 shadow sm:rounded-lg">
        <div className="flex justify-between items-center mb-4">
          <h2 className="text-xl font-semibold text-gray-900">Filter Posts</h2>
          <SortSelect
            availableFields={SORT_FIELDS}
            onChange={setSortFields}
            className="min-w-[200px] text-gray-900"
          />
        </div>
        <div className="grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-3">
          <div>
            <label className="block text-sm font-medium text-gray-700">Date From</label>
            <input
              type="date"
              className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 sm:text-sm text-gray-900"
              value={dateFrom}
              onChange={(e) => setDateFrom(e.target.value)}
            />
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700">Date To</label>
            <input
              type="date"
              className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 sm:text-sm text-gray-900"
              value={dateTo}
              onChange={(e) => setDateTo(e.target.value)}
            />
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700">Visibility</label>
            <select
              className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 sm:text-sm text-gray-900"
              value={visible === undefined ? '' : String(visible)}
              onChange={(e) => setVisible(e.target.value === '' ? undefined : e.target.value === 'true')}
            >
              <option value="">All</option>
              <option value="true">Published</option>
              <option value="false">Draft</option>
            </select>
          </div>
        </div>
      </div>

      {/* Table */}
      <div className="overflow-hidden bg-white shadow sm:rounded-lg">
        <div className="overflow-x-auto">
          <table className="min-w-full divide-y divide-gray-300">
            <thead className="bg-gray-50">
              <tr>
                <th
                  scope="col"
                  className="px-3 py-3.5 text-left text-sm font-semibold text-gray-900"
                >
                  Title
                </th>
                <th
                  scope="col"
                  className="px-3 py-3.5 text-left text-sm font-semibold text-gray-900"
                >
                  Author
                </th>
                <th
                  scope="col"
                  className="px-3 py-3.5 text-left text-sm font-semibold text-gray-900"
                >
                  Published Date
                </th>
                <th
                  scope="col"
                  className="px-3 py-3.5 text-left text-sm font-semibold text-gray-900"
                >
                  Status
                </th>
                <th scope="col" className="relative py-3.5 pl-3 pr-4 sm:pr-6">
                  <span className="sr-only">Actions</span>
                </th>
              </tr>
            </thead>
            <tbody className="divide-y divide-gray-200 bg-white">
              {isLoading ? (
                <tr>
                  <td colSpan={5} className="text-center py-4">Loading...</td>
                </tr>
              ) : data?.Items.map((post) => (
                <tr key={post.id}>
                  <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-900">
                    {post.heading}
                  </td>
                  <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
                    {post.author}
                  </td>
                  <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
                    {new Date(post.publishedDate).toLocaleDateString()}
                  </td>
                  <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
                    <span
                      className={`inline-flex rounded-full px-2 text-xs font-semibold leading-5 ${
                        post.visible
                          ? 'bg-green-100 text-green-800'
                          : 'bg-yellow-100 text-yellow-800'
                      }`}
                    >
                      {post.visible ? 'Published' : 'Draft'}
                    </span>
                  </td>
                  <td className="relative whitespace-nowrap py-4 pl-3 pr-4 text-right text-sm font-medium sm:pr-6">
                    <button
                      onClick={() => {/* TODO: Implement edit */}}
                      className="text-primary-600 hover:text-primary-900 mr-4"
                    >
                      Edit
                    </button>
                    <button
                      onClick={() => toggleVisibilityMutation.mutate(post)}
                      className="text-yellow-600 hover:text-yellow-900 mr-4"
                    >
                      {post.visible ? 'Unpublish' : 'Publish'}
                    </button>
                    <button
                      onClick={() => handleDelete(post.id)}
                      className="text-red-600 hover:text-red-900"
                    >
                      Delete
                    </button>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </div>

      {/* Pagination */}
      {data && data.TotalPages > 0 && (
        <div className="mt-8 flex items-center justify-between bg-white px-4 py-3 sm:px-6">
          <div className="flex flex-1 justify-between sm:flex sm:items-center">
            <div>
              <p className="text-sm text-gray-700">
                Showing <span className="font-medium">{(currentPage - 1) * pageSize + 1}</span> to{' '}
                <span className="font-medium">
                  {Math.min(currentPage * pageSize, data.TotalItems)}
                </span>{' '}
                of <span className="font-medium">{data.TotalItems}</span> results
              </p>
            </div>
            <div className="flex items-center gap-6">
              <div className="flex items-center gap-2">
                <label htmlFor="pageSize" className="text-sm text-gray-600">Show:</label>
                <select
                  id="pageSize"
                  value={pageSize}
                  onChange={(e) => {
                    setPageSize(Number(e.target.value));
                    setCurrentPage(1);
                  }}
                  className="rounded-md border-gray-300 py-1 pl-3 pr-8 text-sm focus:border-primary-500 focus:outline-none focus:ring-primary-500 text-gray-900"
                >
                  <option value="5">5</option>
                  <option value="10">10</option>
                  <option value="20">20</option>
                  <option value="50">50</option>
                </select>
                <span className="text-sm text-gray-600">per page</span>
              </div>
              <Pagination
                currentPage={currentPage}
                totalPages={data.TotalPages}
                onPageChange={setCurrentPage}
              />
            </div>
          </div>
        </div>
      )}
    </div>
  );
};

export default BlogManagement; 
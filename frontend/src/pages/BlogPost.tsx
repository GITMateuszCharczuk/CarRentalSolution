import { useState } from 'react';
import { useParams, Link } from 'react-router-dom';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { useSelector } from 'react-redux';
import { blogService, fileService } from '../services/api';
import { BlogPost as BlogPostType, BlogComment } from '../types/api';
import { selectIsAuthenticated, selectCurrentUser } from '../store/slices/authSlice';

const BlogPost = () => {
  const { id } = useParams<{ id: string }>();
  const queryClient = useQueryClient();
  const isAuthenticated = useSelector(selectIsAuthenticated);
  const user = useSelector(selectCurrentUser);
  const [comment, setComment] = useState('');
  const [currentPage, setCurrentPage] = useState(1);
  const [pageSize] = useState(5);

  const { data: post, isLoading: isPostLoading } = useQuery({
    queryKey: ['blogPost', id],
    queryFn: () => blogService.getBlogPostById(id!),
    enabled: !!id,
  });

  const { data: likesData } = useQuery({
    queryKey: ['blogPostLikes', id],
    queryFn: () => blogService.getBlogPostLikes(id!),
    enabled: !!id,
  });

  const { data: comments, isLoading: areCommentsLoading } = useQuery({
    queryKey: ['blogComments', id, currentPage, pageSize],
    queryFn: () => blogService.getBlogPostComments(id!, {
      current_page: currentPage,
      page_size: pageSize
    }),
    enabled: !!id,
  });

  const addCommentMutation = useMutation({
    mutationFn: (commentData: { description: string }) =>
      blogService.createBlogPostComment(id!, commentData),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['blogComments', id] });
      setComment('');
    },
  });

  const toggleLikeMutation = useMutation({
    mutationFn: () => blogService.likeBlogPost(id!),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['blogPost', id] });
      queryClient.invalidateQueries({ queryKey: ['blogPostLikes', id] });
    },
  });

  const handleCommentSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (!comment.trim()) return;
    addCommentMutation.mutate({ description: comment });
  };

  if (isPostLoading) {
    return (
      <div className="flex h-96 items-center justify-center">
        <div className="h-32 w-32 animate-spin rounded-full border-b-2 border-t-2 border-primary-500"></div>
      </div>
    );
  }

  if (!post?.blog_post) {
    return (
      <div className="flex h-96 items-center justify-center">
        <p className="text-red-500">Blog post not found.</p>
      </div>
    );
  }

  const blogPost = post.blog_post;

  return (
    <div className="mx-auto max-w-4xl space-y-8 px-4 py-8">
      {/* Featured Image */}
      {blogPost.featuredImageUrl && (
        <div className="aspect-h-2 aspect-w-3 overflow-hidden rounded-lg">
          <img
            src={blogPost.featuredImageUrl ? fileService.getFileUrl(blogPost.featuredImageUrl) : '/placeholder-blog.jpg'}
            alt={blogPost.heading}
            className="h-96 w-full object-cover"
          />
        </div>
      )}

      {/* Post Header */}
      <div className="text-center">
        <h1 className="text-4xl font-bold tracking-tight text-gray-900 sm:text-5xl">
          {blogPost.heading}
        </h1>
        <div className="mt-4 flex items-center justify-center space-x-4 text-sm text-gray-500">
          <span>{blogPost.author}</span>
          <span>•</span>
          <time dateTime={blogPost.publishedDate}>
            {new Date(blogPost.publishedDate).toLocaleDateString()}
          </time>
        </div>
        {blogPost.tags && blogPost.tags.length > 0 && (
          <div className="mt-2 flex justify-center space-x-2">
            {blogPost.tags.map((tag: string) => (
              <span
                key={tag}
                className="inline-flex items-center rounded-full bg-primary-100 px-3 py-0.5 text-sm font-medium text-primary-800"
              >
                {tag}
              </span>
            ))}
          </div>
        )}
      </div>

      {/* Post Content */}
      <div className="prose prose-lg mx-auto mt-6 text-black">
        <div dangerouslySetInnerHTML={{ __html: blogPost.content }} />
      </div>

      {/* Like Button */}
      <div className="flex justify-center">
        <button
          onClick={() => isAuthenticated && toggleLikeMutation.mutate()}
          disabled={!isAuthenticated || toggleLikeMutation.isPending}
          className={`inline-flex items-center space-x-2 rounded-full px-4 py-2 text-sm font-medium ${
            isAuthenticated
              ? 'bg-primary-100 text-primary-700 hover:bg-primary-200'
              : 'cursor-not-allowed bg-gray-100 text-gray-500'
          }`}
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            className="h-5 w-5"
            viewBox="0 0 20 20"
            fill="currentColor"
          >
            <path
              fillRule="evenodd"
              d="M3.172 5.172a4 4 0 015.656 0L10 6.343l1.172-1.171a4 4 0 115.656 5.656L10 17.657l-6.828-6.829a4 4 0 010-5.656z"
              clipRule="evenodd"
            />
          </svg>
          <span>{likesData?.TotalCount || 0} likes</span>
        </button>
      </div>

      {/* Comments Section */}
      <div className="mt-8">
        <h2 className="text-2xl font-bold text-gray-900">Comments</h2>

        {/* Comment Form */}
        {isAuthenticated ? (
          <form onSubmit={handleCommentSubmit} className="mt-4">
            <div>
              <label htmlFor="comment" className="sr-only">
                Add your comment
              </label>
              <textarea
                id="comment"
                rows={3}
                className="block w-full rounded-md border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 sm:text-sm"
                placeholder="Add your comment..."
                value={comment}
                onChange={(e) => setComment(e.target.value)}
              />
            </div>
            <div className="mt-3 flex justify-end">
              <button
                type="submit"
                disabled={addCommentMutation.isPending}
                className="inline-flex items-center rounded-md bg-primary-600 px-4 py-2 text-sm font-medium text-white shadow-sm hover:bg-primary-700 focus:outline-none focus:ring-2 focus:ring-primary-500 focus:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
              >
                {addCommentMutation.isPending ? 'Posting...' : 'Post Comment'}
              </button>
            </div>
          </form>
        ) : (
          <p className="mt-4 text-center text-sm text-gray-500">
            Please{' '}
            <Link to="/login" className="text-primary-600 hover:text-primary-500">
              sign in
            </Link>{' '}
            to leave a comment.
          </p>
        )}

        {/* Comments List */}
        <div className="mt-8 space-y-6">
          {areCommentsLoading ? (
            <div className="flex justify-center">
              <div className="h-8 w-8 animate-spin rounded-full border-b-2 border-t-2 border-primary-500"></div>
            </div>
          ) : comments?.Items?.length ? (
            <>
              <div className="space-y-6">
                {comments.Items.map((comment: BlogComment) => (
                  <div key={comment.id} className="flex space-x-4">
                    <div className="flex-shrink-0">
                      <div className="h-10 w-10 rounded-full bg-gray-200"></div>
                    </div>
                    <div className="flex-grow">
                      <div className="flex items-center justify-between">
                        <h3 className="text-sm font-medium text-gray-900">
                          {comment.userId === user?.id ? 'You' : 'Anonymous User'}
                        </h3>
                        <p className="text-sm text-gray-500">
                          {new Date(comment.createdAt).toLocaleDateString()}
                        </p>
                      </div>
                      <p className="mt-1 text-sm text-gray-700">{comment.description}</p>
                    </div>
                  </div>
                ))}
              </div>

              {/* Comments Pagination */}
              {comments.TotalPages > 1 && (
                <div className="mt-6 flex items-center justify-center space-x-2">
                  <button
                    onClick={() => setCurrentPage((prev) => Math.max(prev - 1, 1))}
                    disabled={currentPage === 1}
                    className="rounded-md bg-white px-3 py-2 text-sm font-semibold text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 hover:bg-gray-50 disabled:cursor-not-allowed disabled:opacity-50"
                  >
                    Previous
                  </button>
                  <span className="text-sm text-gray-700">
                    Page {currentPage} of {comments.TotalPages}
                  </span>
                  <button
                    onClick={() => setCurrentPage((prev) => Math.min(prev + 1, comments.TotalPages))}
                    disabled={currentPage === comments.TotalPages}
                    className="rounded-md bg-white px-3 py-2 text-sm font-semibold text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 hover:bg-gray-50 disabled:cursor-not-allowed disabled:opacity-50"
                  >
                    Next
                  </button>
                </div>
              )}
            </>
          ) : (
            <p className="text-center text-gray-500">No comments yet. Be the first to comment!</p>
          )}
        </div>
      </div>
    </div>
  );
};

export default BlogPost; 
import getHandler from '@/handlers/get_handler';
import Toaster from '@/utils/toaster';
import React, { useEffect, useState } from 'react';
import { SERVER_ERROR } from '@/config/errors';
import Loader from '@/components/loader';
import { Post } from '@/types';
import InfiniteScroll from 'react-infinite-scroll-component';
import PostComponent from '@/components/post';

const Posts = () => {
  const [posts, setPosts] = useState<Post[]>([]);
  const [loading, setLoading] = useState(true);
  const [hasMore, setHasMore] = useState(true);
  const [page, setPage] = useState(1);

  const getPosts = () => {
    setLoading(true);
    const URL = `/flags/posts?page=${page}&limit=${10}`;
    getHandler(URL)
      .then(res => {
        if (res.statusCode === 200) {
          const addedPosts = [...posts, ...(res.data.posts || [])];
          if (addedPosts.length === posts.length) setHasMore(false);
          setPosts(addedPosts);
          setPage(prev => prev + 1);
          setLoading(false);
        } else {
          if (res.data.message) Toaster.error(res.data.message, 'error_toaster');
          else {
            Toaster.error(SERVER_ERROR, 'error_toaster');
          }
        }
      })
      .catch(err => {
        Toaster.error(SERVER_ERROR, 'error_toaster');
      });
  };

  useEffect(() => {
    getPosts();
  }, []);

  return (
    <div className="w-full h-full p-4">
      {loading ? (
        <div className="w-[45vw] mx-auto max-lg:w-[85%] max-md:w-full">
          <Loader />
        </div>
      ) : (
        <InfiniteScroll
          className="w-[45vw] mx-auto max-lg:w-[85%] max-md:w-screen flex flex-col gap-2 max-lg:px-4 pb-base_padding"
          dataLength={posts.length}
          next={getPosts}
          hasMore={hasMore}
          loader={<Loader />}
        >
          {posts.length === 0 ? (
            <div>No Flagged Posts.</div>
          ) : (
            posts.map(post => <PostComponent key={post.id} post={post} setPosts={setPosts} />)
          )}
        </InfiniteScroll>
      )}
    </div>
  );
};

export default Posts;

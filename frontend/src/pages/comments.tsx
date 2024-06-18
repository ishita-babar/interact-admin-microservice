import getHandler from '@/handlers/get_handler';
import Toaster from '@/utils/toaster';
import React, { useEffect, useState } from 'react';
import { SERVER_ERROR } from '@/config/errors';
import Loader from '@/components/loader';
import { Comment } from '@/types';
import InfiniteScroll from 'react-infinite-scroll-component';
import CommentComponent from '@/components/comment';
import BaseWrapper from '@/wrappers/base';
import Sidebar from '@/components/common/sidebar';
import MainWrapper from '@/wrappers/main';

const Comments = () => {
  const [comments, setComments] = useState<Comment[]>([]);
  const [loading, setLoading] = useState(true);
  const [hasMore, setHasMore] = useState(true);
  const [page, setPage] = useState(1);

  const getComments = () => {
    setLoading(true);
    const URL = `/flags/comments?page=${page}&limit=${10}`;
    getHandler(URL)
      .then(res => {
        if (res.statusCode === 200) {
          const addedComments = [...comments, ...(res.data.posts || [])];
          if (addedComments.length === comments.length) setHasMore(false);
          setComments(addedComments);
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
    getComments();
  }, []);

  return (
    <BaseWrapper>
      <Sidebar index={10} />
      <MainWrapper>
    <div className="w-full h-full p-4">
      {loading ? (
        <div className="w-[45vw] mx-auto max-lg:w-[85%] max-md:w-full">
          <Loader />
        </div>
      ) : (
        <InfiniteScroll
          className="w-[45vw] mx-auto max-lg:w-[85%] max-md:w-screen flex flex-col gap-2 max-lg:px-4 pb-base_padding"
          dataLength={comments.length}
          next={getComments}
          hasMore={hasMore}
          loader={<Loader />}
        >
          {comments.length === 0 ? (
            <div>No Flagged Comments.</div>
          ) : (
            comments.map(comment => <CommentComponent key={comment.id} comment={comment} setComments={setComments} />)
          )}
        </InfiniteScroll>
      )}
    </div>
    </MainWrapper>
    </BaseWrapper>
  );
};

export default Comments;
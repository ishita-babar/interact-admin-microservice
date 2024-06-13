import getHandler from '@/handlers/get_handler';
import Toaster from '@/utils/toaster';
import React, { useEffect, useState } from 'react';
import { SERVER_ERROR } from '@/config/errors';
import Loader from '@/components/loader';
import { Poll } from '@/types';
import InfiniteScroll from 'react-infinite-scroll-component';
import PollComponent from '@/components/poll';

const Polls = () => {
  const [polls, setPolls] = useState<Poll[]>([]);
  const [loading, setLoading] = useState(true);
  const [hasMore, setHasMore] = useState(true);
  const [page, setPage] = useState(1);

  const getPolls = () => {
    setLoading(true);
    const URL = `/flags/polls?page=${page}&limit=${10}`;
    getHandler(URL)
      .then(res => {
        if (res.statusCode === 200) {
          const addedPolls = [...polls, ...(res.data.polls || [])];
          if (addedPolls.length === polls.length) setHasMore(false);
          setPolls(addedPolls);
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
    getPolls();
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
          dataLength={polls.length}
          next={getPolls}
          hasMore={hasMore}
          loader={<Loader />}
        >
          {polls.length === 0 ? (
            <div>No Flagged Polls.</div>
          ) : (
            polls.map(poll => <PollComponent key={poll.id} poll={poll} setPolls={setPolls} />)
          )}
        </InfiniteScroll>
      )}
    </div>
  );
};

export default Polls;
